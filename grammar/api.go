package grammar

import (
	"context"
	"runtime"
	"strings"
	"sync"
)

// Parser is a tree-sitter parser backed by a wazero instance of a Language.
//
// Ownership is GC-managed. Delete is optional.
// Methods are safe for concurrent use: an internal mutex serializes wasm calls.
// Parse returns a pure-Go snapshot. The wasm tree is freed before return.
type Parser struct {
	state   *parserState
	cleanup runtime.Cleanup
}

// parserState is heap-allocated so runtime.Cleanup can free the wasm instance
// without keeping the *Parser reachable.
type parserState struct {
	mu     sync.Mutex
	lang   *Language
	guest  *guest
	parser uint32
}

func (s *parserState) close() {
	if s == nil {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.closeLocked(context.Background())
}

func (s *parserState) closeLocked(ctx context.Context) {
	if s.guest != nil && s.parser != 0 {
		_, _ = s.guest.call(ctx, "wts_parser_delete", uint64(s.parser))
	}
	s.parser = 0
	if s.guest != nil {
		s.guest.close(ctx)
		s.guest = nil
	}
	s.lang = nil
}

// Tree is an immutable pure-Go snapshot of a parse tree.
type Tree struct {
	root *nodeData
}

// Node is a handle into a Tree snapshot.
// Keep the *Tree reachable while using Nodes from it.
type Node struct {
	data *nodeData
	tree *Tree
}

type nodeData struct {
	typ        string
	start, end uint32
	named      bool
	extra      bool
	isError    bool
	hasError   bool
	hasChanges bool
	fields     []string
	children   []*nodeData
}

// NewParser creates a parser. Callers need not Delete.
func NewParser() *Parser {
	s := &parserState{}
	p := &Parser{state: s}
	p.cleanup = runtime.AddCleanup(p, (*parserState).close, s)
	return p
}

// SetLanguage instantiates lang for this parser.
// A later call replaces the previous instance.
func (p *Parser) SetLanguage(lang *Language) bool {
	if p == nil || lang == nil {
		return false
	}
	p.state.mu.Lock()
	defer p.state.mu.Unlock()
	ctx := context.Background()
	p.state.closeLocked(ctx)
	g, err := instantiate(ctx, lang)
	if err != nil {
		return false
	}
	parser, err := g.call1(ctx, "wts_parser_new")
	if err != nil || parser == 0 {
		g.close(ctx)
		return false
	}
	ok, err := g.call1(ctx, "wts_parser_set_language", uint64(parser))
	if err != nil || ok == 0 {
		_, _ = g.call(ctx, "wts_parser_delete", uint64(parser))
		g.close(ctx)
		return false
	}
	p.state.lang = lang
	p.state.guest = g
	p.state.parser = parser
	return true
}

// ParseString parses source into a snapshot Tree.
func (p *Parser) ParseString(source string) *Tree {
	return p.ParseBytes([]byte(source))
}

// ParseBytes parses a contiguous UTF-8 buffer into a snapshot Tree.
func (p *Parser) ParseBytes(source []byte) *Tree {
	if p == nil {
		return &Tree{}
	}
	p.state.mu.Lock()
	defer p.state.mu.Unlock()
	if p.state.guest == nil || p.state.parser == 0 {
		return &Tree{}
	}
	ctx := context.Background()
	return p.snapshot(ctx, source)
}

func (p *Parser) snapshot(ctx context.Context, source []byte) *Tree {
	g := p.state.guest
	src, err := g.alloc(ctx, uint32(len(source)+1))
	if err != nil || src == 0 {
		return &Tree{}
	}
	defer g.release(ctx, src)
	if len(source) > 0 {
		if !g.mod.Memory().Write(src, source) {
			return &Tree{}
		}
	}
	if !g.mod.Memory().WriteByte(src+uint32(len(source)), 0) {
		return &Tree{}
	}
	tree, err := g.call1(ctx, "wts_parser_parse", uint64(p.state.parser), uint64(src), uint64(len(source)))
	if err != nil || tree == 0 {
		return &Tree{}
	}
	defer func() { _, _ = g.call(ctx, "wts_tree_delete", uint64(tree)) }()

	node, err := g.alloc(ctx, g.nodeSz)
	if err != nil || node == 0 {
		return &Tree{}
	}
	defer g.release(ctx, node)
	if _, err := g.call(ctx, "wts_tree_root", uint64(tree), uint64(node)); err != nil {
		return &Tree{}
	}
	return &Tree{root: captureNode(ctx, g, node)}
}

func captureNode(ctx context.Context, g *guest, node uint32) *nodeData {
	null, err := g.call1(ctx, "wts_node_is_null", uint64(node))
	if err != nil || null != 0 {
		return nil
	}
	typPtr, err := g.call1(ctx, "wts_node_type", uint64(node))
	if err != nil {
		return nil
	}
	count, err := g.call1(ctx, "wts_node_child_count", uint64(node))
	if err != nil {
		return nil
	}
	start, err := g.call1(ctx, "wts_node_start_byte", uint64(node))
	if err != nil {
		return nil
	}
	end, err := g.call1(ctx, "wts_node_end_byte", uint64(node))
	if err != nil {
		return nil
	}
	named, _ := g.call1(ctx, "wts_node_is_named", uint64(node))
	extra, _ := g.call1(ctx, "wts_node_is_extra", uint64(node))
	isErr, _ := g.call1(ctx, "wts_node_is_error", uint64(node))
	hasErr, _ := g.call1(ctx, "wts_node_has_error", uint64(node))
	hasCh, _ := g.call1(ctx, "wts_node_has_changes", uint64(node))
	d := &nodeData{
		typ:        g.cstring(typPtr),
		start:      start,
		end:        end,
		named:      named != 0,
		extra:      extra != 0,
		isError:    isErr != 0,
		hasError:   hasErr != 0,
		hasChanges: hasCh != 0,
	}
	if count == 0 {
		return d
	}
	child, err := g.alloc(ctx, g.nodeSz)
	if err != nil || child == 0 {
		return d
	}
	defer g.release(ctx, child)
	d.children = make([]*nodeData, 0, count)
	d.fields = make([]string, 0, count)
	for i := uint32(0); i < count; i++ {
		fieldPtr, _ := g.call1(ctx, "wts_node_field_name_for_child", uint64(node), uint64(i))
		if _, err := g.call(ctx, "wts_node_child", uint64(node), uint64(i), uint64(child)); err != nil {
			d.fields = append(d.fields, "")
			d.children = append(d.children, nil)
			continue
		}
		d.fields = append(d.fields, g.cstring(fieldPtr))
		d.children = append(d.children, captureNode(ctx, g, child))
	}
	return d
}

// Delete eagerly frees the wasm parser. Optional.
func (p *Parser) Delete() {
	if p == nil || p.state == nil {
		return
	}
	p.cleanup.Stop()
	p.state.close()
}

// Delete drops the snapshot root. Optional. There is no wasm state.
func (t *Tree) Delete() {
	if t == nil {
		return
	}
	t.root = nil
}

// RootNode returns the root node of the tree.
func (t *Tree) RootNode() *Node {
	if t == nil || t.root == nil {
		return &Node{}
	}
	return &Node{data: t.root, tree: t}
}

// Type returns the node type.
func (n *Node) Type() string {
	if n == nil || n.data == nil {
		return ""
	}
	return n.data.typ
}

// ChildCount returns the number of children.
func (n *Node) ChildCount() uint32 {
	if n == nil || n.data == nil {
		return 0
	}
	return uint32(len(n.data.children))
}

// Child returns the child at index.
func (n *Node) Child(index uint32) *Node {
	if n == nil || n.data == nil || int(index) >= len(n.data.children) {
		return &Node{}
	}
	return &Node{data: n.data.children[index], tree: n.tree}
}

// FieldNameForChild returns the field name for the child at index.
func (n *Node) FieldNameForChild(index uint32) string {
	if n == nil || n.data == nil || int(index) >= len(n.data.fields) {
		return ""
	}
	return n.data.fields[index]
}

// NamedChildCount returns the number of named children.
func (n *Node) NamedChildCount() uint32 {
	if n == nil || n.data == nil {
		return 0
	}
	var c uint32
	for _, ch := range n.data.children {
		if ch != nil && ch.named {
			c++
		}
	}
	return c
}

// NamedChild returns the named child at index.
func (n *Node) NamedChild(index uint32) *Node {
	if n == nil || n.data == nil {
		return &Node{}
	}
	var seen uint32
	for _, ch := range n.data.children {
		if ch == nil || !ch.named {
			continue
		}
		if seen == index {
			return &Node{data: ch, tree: n.tree}
		}
		seen++
	}
	return &Node{}
}

// StartByte returns the start byte offset.
func (n *Node) StartByte() uint32 {
	if n == nil || n.data == nil {
		return 0
	}
	return n.data.start
}

// EndByte returns the end byte offset.
func (n *Node) EndByte() uint32 {
	if n == nil || n.data == nil {
		return 0
	}
	return n.data.end
}

// String returns an S-expression of the node.
func (n *Node) String() string {
	if n == nil || n.data == nil {
		return ""
	}
	var b strings.Builder
	writeSexpr(&b, n.data)
	return b.String()
}

func writeSexpr(b *strings.Builder, d *nodeData) {
	if d == nil {
		return
	}
	if !d.named && len(d.children) == 0 {
		b.WriteString(quoteSexprAtom(d.typ))
		return
	}
	b.WriteByte('(')
	b.WriteString(d.typ)
	for i, ch := range d.children {
		if ch == nil {
			continue
		}
		b.WriteByte(' ')
		if i < len(d.fields) && d.fields[i] != "" {
			b.WriteString(d.fields[i])
			b.WriteString(": ")
		}
		writeSexpr(b, ch)
	}
	b.WriteByte(')')
}

func quoteSexprAtom(s string) string {
	return `"` + strings.ReplaceAll(s, `"`, `\"`) + `"`
}

// IsNull reports whether the node is null.
func (n *Node) IsNull() bool {
	return n == nil || n.data == nil
}

// IsNamed reports whether the node is named.
func (n *Node) IsNamed() bool {
	return n != nil && n.data != nil && n.data.named
}

// IsExtra reports whether the node is extra.
func (n *Node) IsExtra() bool {
	return n != nil && n.data != nil && n.data.extra
}

// IsError reports whether the node is an error.
func (n *Node) IsError() bool {
	return n != nil && n.data != nil && n.data.isError
}

// HasError reports whether the node or a descendant is an error.
func (n *Node) HasError() bool {
	return n != nil && n.data != nil && n.data.hasError
}

// HasChanges reports whether the node is marked changed.
func (n *Node) HasChanges() bool {
	return n != nil && n.data != nil && n.data.hasChanges
}

// PrintTree returns the S-expression, or "(null)" when the node is null.
func (n *Node) PrintTree() string {
	if n.IsNull() {
		return "(null)"
	}
	return n.String()
}

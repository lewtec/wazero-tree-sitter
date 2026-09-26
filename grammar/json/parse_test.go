package grammar_json

import (
	"strings"
	"testing"

	"github.com/lewtec/wazero-tree-sitter/grammar"
)

func TestParseObject(t *testing.T) {
	src := []byte(`{"a": 1}`)
	p := grammar.NewParser()
	defer p.Delete()
	if !p.SetLanguage(Language()) {
		t.Fatal("SetLanguage")
	}
	tree := p.ParseBytes(src)
	root := tree.RootNode()
	if root.IsNull() {
		t.Fatal("null root")
	}
	if root.HasError() {
		t.Fatalf("has error: %s", root.String())
	}
	if root.Type() != "document" {
		t.Fatalf("root type %q", root.Type())
	}
	sexpr := root.String()
	if !strings.Contains(sexpr, "pair") || !strings.Contains(sexpr, "number") {
		t.Fatalf("sexpr %s", sexpr)
	}
	node := grammar.BuildParseNode(root, src, "")
	if node == nil || node.Children == nil {
		t.Fatal("empty parse node")
	}
}

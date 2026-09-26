package grammar

import "sort"

// LineIndex maps source byte offsets to 1-based line numbers and 0-based
// byte columns. Columns are byte offsets from the start of the line.
type LineIndex struct {
	starts []int
}

// NewLineIndex builds a line index for text.
func NewLineIndex(text string) *LineIndex {
	return NewLineIndexBytes([]byte(text))
}

// NewLineIndexBytes builds a line index for a byte buffer.
func NewLineIndexBytes(src []byte) *LineIndex {
	starts := make([]int, 1, len(src)/40+2)
	starts[0] = 0
	for i := 0; i < len(src); i++ {
		if src[i] == '\n' && i+1 < len(src) {
			starts = append(starts, i+1)
		}
	}
	return &LineIndex{starts: starts}
}

// LineAt returns the 1-based line containing byte offset off.
func (li *LineIndex) LineAt(off int) int {
	line, _ := li.LineColumnAt(off)
	return line
}

// LineAtU32 is LineAt for tree-sitter byte offsets.
func (li *LineIndex) LineAtU32(off uint32) int {
	return li.LineAt(int(off))
}

// LineColumnAt returns the 1-based line and 0-based byte column for off.
func (li *LineIndex) LineColumnAt(off int) (line, column int) {
	if li == nil || len(li.starts) == 0 {
		return 1, 0
	}
	if off < 0 {
		return 1, 0
	}
	i := sort.Search(len(li.starts), func(i int) bool {
		return li.starts[i] > off
	}) - 1
	if i < 0 {
		return 1, 0
	}
	return i + 1, off - li.starts[i]
}

// LineColumnAtU32 is LineColumnAt for tree-sitter byte offsets.
func (li *LineIndex) LineColumnAtU32(off uint32) (line, column int) {
	return li.LineColumnAt(int(off))
}

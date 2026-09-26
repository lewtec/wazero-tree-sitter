package grammar

import "testing"

func TestLineColumnAt(t *testing.T) {
	li := NewLineIndex("ab\nc")
	line, col := li.LineColumnAt(0)
	if line != 1 || col != 0 {
		t.Fatalf("start: got %d:%d", line, col)
	}
	line, col = li.LineColumnAt(3)
	if line != 2 || col != 0 {
		t.Fatalf("second line: got %d:%d", line, col)
	}
	if li.LineAt(-1) != 1 {
		t.Fatal("negative offset")
	}
}

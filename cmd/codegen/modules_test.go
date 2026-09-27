package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteLangGoModRequiresPublishedCore(t *testing.T) {
	dir := t.TempDir()
	langDir := filepath.Join(dir, "json")
	if err := os.MkdirAll(langDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(langDir, "grammar.wasm"), []byte("wasm"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := writeLangGoMod(dir, "json"); err != nil {
		t.Fatal(err)
	}
	coreVer, err := coreGrammarPseudoVersion()
	if err != nil {
		t.Fatal(err)
	}
	if coreVer == "v0.0.0" || !strings.Contains(coreVer, "-") {
		t.Fatalf("core grammar version must be a commit pseudo-version, got %s", coreVer)
	}
	data, err := os.ReadFile(filepath.Join(langDir, "go.mod"))
	if err != nil {
		t.Fatal(err)
	}
	s := string(data)
	for _, want := range []string{
		"module github.com/lewtec/wazero-tree-sitter/grammar/json",
		"require github.com/lewtec/wazero-tree-sitter/grammar " + coreVer,

	} {
		if !strings.Contains(s, want) {
			t.Errorf("go.mod missing %q\n%s", want, s)
		}
	}
}

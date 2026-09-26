package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// tree-sitter core is placed by core:place at third-party/tree-sitter.
func resolveTreeSitterPath() (string, error) {
	if p := strings.TrimSpace(os.Getenv("TREE_SITTER_PATH")); p != "" {
		return checkTreeSitter(p)
	}
	marker, err := findUp(filepath.Join("third-party", "tree-sitter", "lib", "src", "lib.c"))
	if err != nil {
		return "", fmt.Errorf("placed tree-sitter core: %w (run: workspaced codebase apply)", err)
	}
	return checkTreeSitter(filepath.Dir(filepath.Dir(filepath.Dir(marker))))
}

func checkTreeSitter(root string) (string, error) {
	mark := filepath.Join(root, "lib", "src", "lib.c")
	if _, err := os.Stat(mark); err != nil {
		return "", fmt.Errorf("tree-sitter root %s: %w (need lib/src/lib.c)", root, err)
	}
	return root, nil
}

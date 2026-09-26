package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/lewtec/wazero-tree-sitter/grammar"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:   "parse <language> <file>",
		Short: "Parse a source file and print the tree as JSON",
		Args:  cobra.ExactArgs(2),
		RunE: func(_ *cobra.Command, args []string) error {
			return run(args[0], args[1])
		},
	}
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(languageName, filename string) error {
	source, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	lang, ok := grammar.Get(strings.ToLower(languageName))
	if !ok || lang == nil {
		return fmt.Errorf("unsupported language: %s\nsupported: %s", languageName, grammar.SupportedLanguages())
	}
	parser := grammar.NewParser()
	defer parser.Delete()
	if !parser.SetLanguage(lang) {
		return fmt.Errorf("failed to set language %s", languageName)
	}
	tree := parser.ParseBytes(source)
	root := tree.RootNode()
	if root.IsNull() {
		return fmt.Errorf("parse failed: null root")
	}
	out := grammar.ParseOutput{
		Language: strings.ToLower(languageName),
		File:     filename,
		Root:     grammar.BuildParseNode(root, source, ""),
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(out)
}

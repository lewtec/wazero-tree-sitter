package main

import (
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// GrammarUnit is one compile target: a directory containing src/parser.c.
type GrammarUnit struct {
	Name     string
	Path     string
	ParserC  string
	Priority int
}

var langEntryRe = regexp.MustCompile(`(?m)(?:TS_PUBLIC\s+)?const\s+TSLanguage\s*\*\s*tree_sitter_(\w+)\s*\(\s*void\s*\)`)

func discoverGrammarUnits(thirdPartyGlob string) ([]GrammarUnit, error) {
	repos, err := filepath.Glob(thirdPartyGlob)
	if err != nil {
		return nil, err
	}
	var candidates []GrammarUnit
	for _, repo := range repos {
		info, err := os.Stat(repo)
		if err != nil || !info.IsDir() {
			continue
		}
		err = filepath.WalkDir(repo, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				switch d.Name() {
				case "node_modules", ".git", "target", "build", "prebuilds", ".build":
					return filepath.SkipDir
				}
				return nil
			}
			if d.Name() != "parser.c" || filepath.Base(filepath.Dir(path)) != "src" {
				return nil
			}
			unitPath := filepath.Dir(filepath.Dir(path))
			name := languageNameForParser(path, unitPath)
			if name == "" {
				return nil
			}
			candidates = append(candidates, GrammarUnit{
				Name:     name,
				Path:     unitPath,
				ParserC:  path,
				Priority: grammarPriority(path, repo),
			})
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	sort.Slice(candidates, func(i, j int) bool {
		if candidates[i].Priority != candidates[j].Priority {
			return candidates[i].Priority < candidates[j].Priority
		}
		if candidates[i].Name != candidates[j].Name {
			return candidates[i].Name < candidates[j].Name
		}
		return candidates[i].Path < candidates[j].Path
	})
	seen := map[string]struct{}{}
	out := make([]GrammarUnit, 0, len(candidates))
	for _, u := range candidates {
		if _, ok := seen[u.Name]; ok {
			continue
		}
		seen[u.Name] = struct{}{}
		out = append(out, u)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out, nil
}

func languageNameForParser(parserC, unitPath string) string {
	data, err := os.ReadFile(parserC)
	if err != nil {
		if !os.IsNotExist(err) {
			slog.Warn("failed to read parser.c", "path", parserC, "error", err)
		}
		return normalizeGrammarName(filepath.Base(unitPath))
	}
	if m := langEntryRe.FindSubmatch(data); len(m) == 2 {
		return string(m[1])
	}
	return normalizeGrammarName(filepath.Base(unitPath))
}

func normalizeGrammarName(name string) string {
	name = filepath.Base(name)
	const prefix = "tree-sitter-"
	if strings.HasPrefix(name, prefix) {
		name = name[len(prefix):]
	}
	return strings.ReplaceAll(name, "-", "_")
}

func grammarPriority(parserC, repoRoot string) int {
	rel, err := filepath.Rel(repoRoot, parserC)
	if err != nil {
		rel = parserC
	}
	parts := strings.Split(filepath.ToSlash(rel), "/")
	score := 0
	for _, p := range parts {
		switch p {
		case "examples", "test", "tests", "corpus":
			score += 100
		case "schema":
			score += 50
		case "dialects", "dialect":
			score += 40
		}
	}
	if depth := len(parts) - 1; depth > 0 {
		score += depth
	}
	return score
}

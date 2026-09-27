package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"golang.org/x/mod/modfile"
	"golang.org/x/mod/module"
)

const (
	rootModulePath    = "github.com/lewtec/wazero-tree-sitter"
	grammarModulePath = rootModulePath + "/grammar"
	moduleGoVersion   = "1.27.0"
	localPseudoVer    = "v0.0.0"
	wazeroModulePath  = "github.com/tetratelabs/wazero"
	wazeroModuleVer   = "v1.9.0"
)

func ensureGrammarModules(outputDir string) error {
	abs, err := filepath.Abs(outputDir)
	if err != nil {
		return err
	}
	grammarDir := filepath.Join(abs, "grammar")
	if err := os.MkdirAll(grammarDir, 0o755); err != nil {
		return err
	}
	if err := writeCoreGoMod(grammarDir); err != nil {
		return err
	}
	langs, err := listGrammarLangs(grammarDir)
	if err != nil {
		return err
	}
	for _, lang := range langs {
		if err := writeLangGoMod(grammarDir, lang); err != nil {
			return err
		}
	}
	if err := writeGoWork(abs, langs); err != nil {
		return err
	}
	return updateRootGoMod(abs, langs)
}

func listGrammarLangs(grammarDir string) ([]string, error) {
	entries, err := os.ReadDir(grammarDir)
	if err != nil {
		return nil, err
	}
	var langs []string
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		name := e.Name()
		if _, err := os.Stat(filepath.Join(grammarDir, name, "grammar.wasm")); err != nil {
			continue
		}
		langs = append(langs, name)
	}
	sort.Strings(langs)
	return langs, nil
}

func writeCoreGoMod(grammarDir string) error {
	content := fmt.Sprintf(`module %s

go %s

require %s %s
`, grammarModulePath, moduleGoVersion, wazeroModulePath, wazeroModuleVer)
	return os.WriteFile(filepath.Join(grammarDir, "go.mod"), []byte(content), 0o644)
}

// coreGrammarPseudoVersion is the core grammar module at HEAD.
// Consumers ignore replace directives. A require of v0.0.0 sorts above every
// v0.0.0-* commit, so minimal version selection would demand a tag that does
// not exist. A real pseudo-version is fetchable without a consumer replace.
func coreGrammarPseudoVersion() (string, error) {
	rev, err := gitOutput("rev-parse", "HEAD")
	if err != nil {
		return "", err
	}
	if len(rev) < 12 {
		return "", fmt.Errorf("git HEAD %q is too short for a pseudo-version", rev)
	}
	unixStr, err := gitOutput("log", "-1", "--format=%ct", "HEAD")
	if err != nil {
		return "", err
	}
	unix, err := strconv.ParseInt(unixStr, 10, 64)
	if err != nil {
		return "", fmt.Errorf("parse HEAD commit time %q: %w", unixStr, err)
	}
	return module.PseudoVersion("v0", "", time.Unix(unix, 0).UTC(), rev[:12]), nil
}

func gitOutput(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("git %s: %w", strings.Join(args, " "), err)
	}
	return strings.TrimSpace(string(out)), nil
}

func writeLangGoMod(grammarDir, lang string) error {
	coreVer, err := coreGrammarPseudoVersion()
	if err != nil {
		return err
	}
	content := fmt.Sprintf(`module %s/%s

go %s

require %s %s
`, grammarModulePath, lang, moduleGoVersion, grammarModulePath, coreVer)
	return os.WriteFile(filepath.Join(grammarDir, lang, "go.mod"), []byte(content), 0o644)
}

func writeGoWork(outputDir string, langs []string) error {
	var b strings.Builder
	b.WriteString("go " + moduleGoVersion + "\n\nuse (\n\t.\n\t./grammar\n")
	for _, lang := range langs {
		fmt.Fprintf(&b, "\t./grammar/%s\n", lang)
	}
	b.WriteString(")\n")
	return os.WriteFile(filepath.Join(outputDir, "go.work"), []byte(b.String()), 0o644)
}

func updateRootGoMod(outputDir string, langs []string) error {
	path := filepath.Join(outputDir, "go.mod")
	data, err := os.ReadFile(path)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		data = []byte("module " + rootModulePath + "\n\ngo " + moduleGoVersion + "\n")
	}
	f, err := modfile.Parse(path, data, nil)
	if err != nil {
		return err
	}
	isManaged := func(p string) bool {
		return p == grammarModulePath || strings.HasPrefix(p, grammarModulePath+"/")
	}
	for {
		var drop string
		for _, r := range f.Require {
			if isManaged(r.Mod.Path) {
				drop = r.Mod.Path
				break
			}
		}
		if drop == "" {
			break
		}
		if err := f.DropRequire(drop); err != nil {
			return err
		}
	}
	for {
		var drop string
		for _, r := range f.Replace {
			if isManaged(r.Old.Path) {
				drop = r.Old.Path
				break
			}
		}
		if drop == "" {
			break
		}
		if err := f.DropReplace(drop, ""); err != nil {
			return err
		}
	}
	coreVer, err := coreGrammarPseudoVersion()
	if err != nil {
		return err
	}
	mods := []string{grammarModulePath}
	for _, lang := range langs {
		mods = append(mods, grammarModulePath+"/"+lang)
	}
	for _, m := range mods {
		if err := f.AddRequire(m, coreVer); err != nil {
			return err
		}
	}
	f.Cleanup()
	out, err := f.Format()
	if err != nil {
		return err
	}
	return os.WriteFile(path, out, 0o644)
}

func tidyModules(outputDir string) error {
	if err := execGo(outputDir, "work", "sync"); err != nil {
		return err
	}
	return execGo(outputDir, "mod", "tidy")
}

func execGo(dir string, args ...string) error {
	cmd := exec.Command("go", args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

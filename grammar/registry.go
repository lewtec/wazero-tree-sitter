package grammar

import (
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

var (
	registry = make(map[string]*Language)
	mu       sync.RWMutex
)

// extensionToLanguage maps file extensions (no leading dot, lowercase) to
// registered language names. Extensions that equal a registered name are
// handled by Get first.
var extensionToLanguage = map[string]string{
	"js":       "javascript",
	"mjs":      "javascript",
	"cjs":      "javascript",
	"jsx":      "javascript",
	"ts":       "typescript",
	"mts":      "typescript",
	"cts":      "typescript",
	"tsx":      "tsx",
	"py":       "python",
	"pyi":      "python",
	"rb":       "ruby",
	"rs":       "rust",
	"go":       "go",
	"c":        "c",
	"h":        "c",
	"cc":       "cpp",
	"cpp":      "cpp",
	"cxx":      "cpp",
	"hpp":      "cpp",
	"hh":       "cpp",
	"hxx":      "cpp",
	"lua":      "lua",
	"json":     "json",
	"css":      "css",
	"html":     "html",
	"htm":      "html",
	"xml":      "xml",
	"java":     "java",
	"kt":       "kotlin",
	"kts":      "kotlin",
	"scala":    "scala",
	"cs":       "c_sharp",
	"rspec":    "ruby",
	"sh":       "bash",
	"bash":     "bash",
	"yaml":     "yaml",
	"yml":      "yaml",
	"toml":     "toml",
	"md":       "markdown",
	"markdown": "markdown",
	"nix":      "nix",
	"tf":       "terraform",
	"hcl":      "hcl",
	"diff":     "diff",
	"patch":    "diff",
}

// Register associates name with lang. Re-registering overwrites.
// An empty name is stored as a normal key if passed.
func Register(name string, lang *Language) {
	mu.Lock()
	defer mu.Unlock()
	registry[name] = lang
}

// Get returns the language registered under name.
// Lookup is case-sensitive.
func Get(name string) (*Language, bool) {
	mu.RLock()
	defer mu.RUnlock()
	lang, ok := registry[name]
	return lang, ok
}

// List returns registered grammar names in sorted order.
func List() []string {
	mu.RLock()
	defer mu.RUnlock()
	names := make([]string, 0, len(registry))
	for name := range registry {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// GetByExtension finds a grammar from a file name.
// Matching is case-insensitive on the extension.
func GetByExtension(filename string) (*Language, bool) {
	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(filename), "."))
	if ext == "" {
		return nil, false
	}
	if lang, ok := Get(ext); ok {
		return lang, ok
	}
	if name, ok := extensionToLanguage[ext]; ok {
		return Get(name)
	}
	return nil, false
}

// SupportedLanguages returns registered names joined by ", ", or "none".
func SupportedLanguages() string {
	langs := List()
	if len(langs) == 0 {
		return "none"
	}
	return strings.Join(langs, ", ")
}

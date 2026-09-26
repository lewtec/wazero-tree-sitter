# wazero-tree-sitter

Tree-sitter grammars compiled with emscripten and run with wazero. No CGO.

The layout matches [leaven-tree-sitter](https://github.com/lewtec/leaven-tree-sitter) and [ccgo-tree-sitter](https://github.com/modernc-tree-sitter/ccgo-tree-sitter): one Go module for the host API (`grammar`) and one nested module per language (`grammar/<lang>`). Blank-import a language package so its `init` calls `grammar.Register`.

Each language wasm is self-contained. It links tree-sitter's C library, that grammar's `parser.c` (and `scanner.c` when present), and `runtime/shim.c`. The shim exports integer-only functions. The host copies the tree into a Go snapshot and frees the wasm tree before `Parse` returns.

```bash
mise install
mise run grammars:lock    # go tool modot mod lock
mise run grammars:sync    # go tool modot codebase apply
mise run codegen          # go run ./cmd/codegen
mise run test
```

Uncomment a `#grammar` block in `modot.cue`, then lock, sync, and codegen again.

```go
import (
    "github.com/lewtec/wazero-tree-sitter/grammar"
    _ "github.com/lewtec/wazero-tree-sitter/grammar/json"
)

parser := grammar.NewParser()
lang, _ := grammar.Get("json")
parser.SetLanguage(lang)
tree := parser.ParseString(`{"a": 1}`)
tree.RootNode().String()
```

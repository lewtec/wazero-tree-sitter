// Package grammar binds tree-sitter through wazero (no CGO).
//
// Each language is a self-contained wasm module: tree-sitter core, that
// grammar's parser.c (and scanner, if it has one), and runtime/shim.c.
// mise conda:emscripten builds the module. Codegen embeds it in
// grammar/<lang>.
//
// # Ownership
//
// A Language is the immutable wasm image. Parsers instantiate it. Parse
// copies the syntax tree into a pure-Go snapshot and deletes the wasm tree
// before returning. Delete on Parser or Tree is optional.
//
// # Concurrency
//
// Language values may be shared. Each Parser has its own wasm instance.
// Parser methods serialize on an internal mutex. Tree and Node methods only
// read the snapshot.
//
// # Registry
//
// Blank-import a language package so its init calls Register:
//
//	import _ "github.com/lewtec/wazero-tree-sitter/grammar/json"
//
// Line and column lookup is github.com/lewtec/lewkit/x/text.LineIndex.
//
// # Generated files
//
// grammar/<lang>/api.go and grammar.wasm are codegen output
// (mise run codegen). Do not edit them by hand.
package grammar

module github.com/lewtec/wazero-tree-sitter

go 1.27.0

require (
	github.com/lewtec/wazero-tree-sitter/grammar v0.0.0
	github.com/lewtec/wazero-tree-sitter/grammar/json v0.0.0
	github.com/spf13/cobra v1.10.2
	golang.org/x/mod v0.41.0
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.10 // indirect
	github.com/tetratelabs/wazero v1.12.0 // indirect
	golang.org/x/sys v0.48.0 // indirect
)

replace github.com/lewtec/wazero-tree-sitter/grammar => ./grammar

replace github.com/lewtec/wazero-tree-sitter/grammar/json => ./grammar/json

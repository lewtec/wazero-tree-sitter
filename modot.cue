package modot

// One declaration per language (CUE unifies into #grammar).
// Live #grammar entries are enabled. Uncomment a block to unlock.
// Then: mise run grammars:lock && mise run grammars:sync && mise run codegen
//
// Day-one live grammar is json (parser.c only, no external scanner).

#grammar: json: {
	from: "github:tree-sitter/tree-sitter-json"
	repo: "tree-sitter-json"
}

// #grammar: go: {
// 	from: "github:tree-sitter/tree-sitter-go"
// 	repo: "tree-sitter-go"
// }
//
// #grammar: python: {
// 	from: "github:tree-sitter/tree-sitter-python"
// 	repo: "tree-sitter-python"
// }

#grammar: [string]: {
	from:    string
	version: string | *"HEAD"
	paths:   [...string] | *["src"]
	// Dest basename under third-party/ (usually tree-sitter-<key>).
	repo: string
}

// Core C library (lib/src, lib/include). Source + lock only — not core:place'd.
// Resolve: mise run tree-sitter:path  (or TREE_SITTER_PATH).
#tree_sitter: {
	from:    "github:tree-sitter/tree-sitter"
	version: "HEAD"
}

inputs: {
	tree_sitter: {
		from:    #tree_sitter.from
		version: #tree_sitter.version
	}
	for name, g in #grammar {
		"grammar_\(name)": {
			from:    g.from
			version: g.version
		}
	}
}
modules: {
	// tree_sitter has no place module. Each #grammar is copied under third-party/.
	for name, g in #grammar {
		"grammar_\(name)": {
			from: "core:place"
			config: {
				ignore_missing: true
				items: {
					for p in g.paths {
						"third-party/\(g.repo)/\(p)": "grammar_\(name):\(p)"
					}
				}
			}
		}
	}
}

package workspaced

// Grammars match leaven-tree-sitter. Uncommented #grammar entries are live.
// Then: mise run grammars:lock && mise run grammars:sync && mise run codegen
//
// Pinned to workspaced 0.14.3, the same release leaven-tree-sitter uses.

#grammar: fortran: {
	from: "github:stadelmanma/tree-sitter-fortran"
	repo: "tree-sitter-fortran"
}

#grammar: ada: {
	from: "github:briot/tree-sitter-ada"
	repo: "tree-sitter-ada"
}

#grammar: agda: {
	from: "github:tree-sitter/tree-sitter-agda"
	repo: "tree-sitter-agda"
}

#grammar: angular: {
	from: "github:dlvandenberg/tree-sitter-angular"
	repo: "tree-sitter-angular"
}

#grammar: arduino: {
	from: "github:tree-sitter-grammars/tree-sitter-arduino"
	repo: "tree-sitter-arduino"
}

#grammar: asm: {
	from: "github:RubixDev/tree-sitter-asm"
	repo: "tree-sitter-asm"
}

#grammar: astro: {
	from: "github:virchau13/tree-sitter-astro"
	repo: "tree-sitter-astro"
}

#grammar: bash: {
	from: "github:tree-sitter/tree-sitter-bash"
	repo: "tree-sitter-bash"
}

#grammar: beancount: {
	from: "github:polarmutex/tree-sitter-beancount"
	repo: "tree-sitter-beancount"
}

#grammar: bibtex: {
	from: "github:latex-lsp/tree-sitter-bibtex"
	repo: "tree-sitter-bibtex"
}

#grammar: bicep: {
	from: "github:tree-sitter-grammars/tree-sitter-bicep"
	repo: "tree-sitter-bicep"
}

#grammar: bitbake: {
	from: "github:tree-sitter-grammars/tree-sitter-bitbake"
	repo: "tree-sitter-bitbake"
}

#grammar: blade: {
	from: "github:EmranMR/tree-sitter-blade"
	repo: "tree-sitter-blade"
}

#grammar: c: {
	from: "github:tree-sitter/tree-sitter-c"
	repo: "tree-sitter-c"
}

#grammar: "c-sharp": {
	from: "github:tree-sitter/tree-sitter-c-sharp"
	repo: "tree-sitter-c-sharp"
}

#grammar: cairo: {
	from: "github:tree-sitter-grammars/tree-sitter-cairo"
	repo: "tree-sitter-cairo"
}

#grammar: capnp: {
	from: "github:tree-sitter-grammars/tree-sitter-capnp"
	repo: "tree-sitter-capnp"
}

#grammar: chatito: {
	from: "github:tree-sitter-grammars/tree-sitter-chatito"
	repo: "tree-sitter-chatito"
}

#grammar: clojure: {
	from: "github:sogaiu/tree-sitter-clojure"
	repo: "tree-sitter-clojure"
}

#grammar: cmake: {
	from: "github:uyha/tree-sitter-cmake"
	repo: "tree-sitter-cmake"
}

#grammar: cobol: {
	from: "github:yutaro-sakamoto/tree-sitter-cobol"
	repo: "tree-sitter-cobol"
}

#grammar: comment: {
	from: "github:stsewd/tree-sitter-comment"
	repo: "tree-sitter-comment"
}

#grammar: commonlisp: {
	from: "github:theHamsta/tree-sitter-commonlisp"
	repo: "tree-sitter-commonlisp"
}

#grammar: cpon: {
	from: "github:tree-sitter-grammars/tree-sitter-cpon"
	repo: "tree-sitter-cpon"
}

#grammar: cpp: {
	from: "github:tree-sitter/tree-sitter-cpp"
	repo: "tree-sitter-cpp"
}

#grammar: crystal: {
	from: "github:crystal-lang-tools/tree-sitter-crystal"
	repo: "tree-sitter-crystal"
}

#grammar: css: {
	from: "github:tree-sitter/tree-sitter-css"
	repo: "tree-sitter-css"
}

#grammar: cst: {
	from: "github:tree-sitter-grammars/tree-sitter-cst"
	repo: "tree-sitter-cst"
}

#grammar: csv: {
	from: "github:tree-sitter-grammars/tree-sitter-csv"
	repo: "tree-sitter-csv"
	paths: ["csv/src", "tsv/src", "psv/src", "common"]
}

#grammar: cuda: {
	from: "github:tree-sitter-grammars/tree-sitter-cuda"
	repo: "tree-sitter-cuda"
}

#grammar: cue: {
	from: "github:eonpatapon/tree-sitter-cue"
	repo: "tree-sitter-cue"
}

#grammar: cyberchef: {
	from: "github:tree-sitter-grammars/tree-sitter-cyberchef"
	repo: "tree-sitter-cyberchef"
}

#grammar: dart: {
	from: "github:UserNobody14/tree-sitter-dart"
	repo: "tree-sitter-dart"
}

#grammar: diff: {
	from: "github:tree-sitter-grammars/tree-sitter-diff"
	repo: "tree-sitter-diff"
}

#grammar: dockerfile: {
	from: "github:camdencheek/tree-sitter-dockerfile"
	repo: "tree-sitter-dockerfile"
}

#grammar: doxygen: {
	from: "github:tree-sitter-grammars/tree-sitter-doxygen"
	repo: "tree-sitter-doxygen"
}

#grammar: eex: {
	from: "github:connorlay/tree-sitter-eex"
	repo: "tree-sitter-eex"
}

#grammar: elixir: {
	from: "github:elixir-lang/tree-sitter-elixir"
	repo: "tree-sitter-elixir"
}

#grammar: elm: {
	from: "github:elm-tooling/tree-sitter-elm"
	repo: "tree-sitter-elm"
}

#grammar: "embedded-template": {
	from: "github:tree-sitter/tree-sitter-embedded-template"
	repo: "tree-sitter-embedded-template"
}

#grammar: erlang: {
	from: "github:WhatsApp/tree-sitter-erlang"
	repo: "tree-sitter-erlang"
}

#grammar: firrtl: {
	from: "github:tree-sitter-grammars/tree-sitter-firrtl"
	repo: "tree-sitter-firrtl"
}

#grammar: fish: {
	from: "github:ram02z/tree-sitter-fish"
	repo: "tree-sitter-fish"
}

#grammar: fluent: {
	from: "github:tree-sitter/tree-sitter-fluent"
	repo: "tree-sitter-fluent"
}

#grammar: fsharp: {
	from: "github:ionide/tree-sitter-fsharp"
	repo: "tree-sitter-fsharp"
	paths: ["fsharp/src", "fsharp_signature/src", "common"]
}

#grammar: func: {
	from: "github:tree-sitter-grammars/tree-sitter-func"
	repo: "tree-sitter-func"
}

#grammar: gdscript: {
	from: "github:PrestonKnopp/tree-sitter-gdscript"
	repo: "tree-sitter-gdscript"
}

#grammar: "git-config": {
	from: "github:the-mikedavis/tree-sitter-git-config"
	repo: "tree-sitter-git-config"
}

#grammar: "git-rebase": {
	from: "github:the-mikedavis/tree-sitter-git-rebase"
	repo: "tree-sitter-git-rebase"
}

#grammar: gitattributes: {
	from: "github:tree-sitter-grammars/tree-sitter-gitattributes"
	repo: "tree-sitter-gitattributes"
}

#grammar: gitcommit: {
	from: "github:gbprod/tree-sitter-gitcommit"
	repo: "tree-sitter-gitcommit"
}

#grammar: gitignore: {
	from: "github:shunsambongi/tree-sitter-gitignore"
	repo: "tree-sitter-gitignore"
}

#grammar: gleam: {
	from: "github:gleam-lang/tree-sitter-gleam"
	repo: "tree-sitter-gleam"
}

#grammar: glsl: {
	from: "github:tree-sitter-grammars/tree-sitter-glsl"
	repo: "tree-sitter-glsl"
}

#grammar: gn: {
	from: "github:tree-sitter-grammars/tree-sitter-gn"
	repo: "tree-sitter-gn"
}

#grammar: go: {
	from: "github:tree-sitter/tree-sitter-go"
	repo: "tree-sitter-go"
	astdump: {
		root_type: "source_file"
		first_child_type: "package_clause"
	}
}

#grammar: "go-sum": {
	from: "github:tree-sitter-grammars/tree-sitter-go-sum"
	repo: "tree-sitter-go-sum"
}

#grammar: "godot-resource": {
	from: "github:PrestonKnopp/tree-sitter-godot-resource"
	repo: "tree-sitter-godot-resource"
}

#grammar: "gpg-config": {
	from: "github:tree-sitter-grammars/tree-sitter-gpg-config"
	repo: "tree-sitter-gpg-config"
}

#grammar: graph: {
	from: "github:tree-sitter/tree-sitter-graph"
	repo: "tree-sitter-graph"
}

#grammar: graphql: {
	from: "github:bkegley/tree-sitter-graphql"
	repo: "tree-sitter-graphql"
}

#grammar: groovy: {
	from: "github:murtaza64/tree-sitter-groovy"
	repo: "tree-sitter-groovy"
}

#grammar: gstlaunch: {
	from: "github:tree-sitter-grammars/tree-sitter-gstlaunch"
	repo: "tree-sitter-gstlaunch"
}

#grammar: hare: {
	from: "github:tree-sitter-grammars/tree-sitter-hare"
	repo: "tree-sitter-hare"
}

#grammar: haskell: {
	from: "github:tree-sitter/tree-sitter-haskell"
	repo: "tree-sitter-haskell"
}

#grammar: hcl: {
	from: "github:tree-sitter-grammars/tree-sitter-hcl"
	repo: "tree-sitter-hcl"
	paths: ["src", "dialects/terraform/src"]
}

#grammar: heex: {
	from: "github:phoenixframework/tree-sitter-heex"
	repo: "tree-sitter-heex"
}

#grammar: hlsl: {
	from: "github:tree-sitter-grammars/tree-sitter-hlsl"
	repo: "tree-sitter-hlsl"
}

#grammar: html: {
	from: "github:tree-sitter/tree-sitter-html"
	repo: "tree-sitter-html"
}

#grammar: http: {
	from: "github:rest-nvim/tree-sitter-http"
	repo: "tree-sitter-http"
}

#grammar: hyprlang: {
	from: "github:tree-sitter-grammars/tree-sitter-hyprlang"
	repo: "tree-sitter-hyprlang"
}

#grammar: idris: {
	from: "github:kayhide/tree-sitter-idris"
	repo: "tree-sitter-idris"
}

#grammar: ini: {
	from: "github:justinmk/tree-sitter-ini"
	repo: "tree-sitter-ini"
}

#grammar: ispc: {
	from: "github:tree-sitter-grammars/tree-sitter-ispc"
	repo: "tree-sitter-ispc"
}

#grammar: java: {
	from: "github:tree-sitter/tree-sitter-java"
	repo: "tree-sitter-java"
}

#grammar: javascript: {
	from: "github:tree-sitter/tree-sitter-javascript"
	repo: "tree-sitter-javascript"
}

#grammar: jsdoc: {
	from: "github:tree-sitter/tree-sitter-jsdoc"
	repo: "tree-sitter-jsdoc"
}

#grammar: json: {
	from: "github:tree-sitter/tree-sitter-json"
	repo: "tree-sitter-json"
	astdump: {
		root_type: "document"
		first_child_type: "object"
	}
}

#grammar: json5: {
	from: "github:Joakker/tree-sitter-json5"
	repo: "tree-sitter-json5"
}

#grammar: julia: {
	from: "github:tree-sitter/tree-sitter-julia"
	repo: "tree-sitter-julia"
}

#grammar: just: {
	from: "github:casey/tree-sitter-just"
	repo: "tree-sitter-just"
}

#grammar: kconfig: {
	from: "github:tree-sitter-grammars/tree-sitter-kconfig"
	repo: "tree-sitter-kconfig"
}

#grammar: kdl: {
	from: "github:tree-sitter-grammars/tree-sitter-kdl"
	repo: "tree-sitter-kdl"
}

#grammar: kotlin: {
	from: "github:fwcd/tree-sitter-kotlin"
	repo: "tree-sitter-kotlin"
}

#grammar: latex: {
	from: "github:latex-lsp/tree-sitter-latex"
	repo: "tree-sitter-latex"
}

#grammar: linkerscript: {
	from: "github:tree-sitter-grammars/tree-sitter-linkerscript"
	repo: "tree-sitter-linkerscript"
}

#grammar: llvm: {
	from: "github:benwilliamgraham/tree-sitter-llvm"
	repo: "tree-sitter-llvm"
}

#grammar: lua: {
	from: "github:tree-sitter-grammars/tree-sitter-lua"
	repo: "tree-sitter-lua"
	astdump: {
		root_type: "chunk"
	}
}

#grammar: luadoc: {
	from: "github:tree-sitter-grammars/tree-sitter-luadoc"
	repo: "tree-sitter-luadoc"
}

#grammar: luap: {
	from: "github:tree-sitter-grammars/tree-sitter-luap"
	repo: "tree-sitter-luap"
}

#grammar: luau: {
	from: "github:tree-sitter-grammars/tree-sitter-luau"
	repo: "tree-sitter-luau"
}

#grammar: make: {
	from: "github:tree-sitter-grammars/tree-sitter-make"
	repo: "tree-sitter-make"
}

#grammar: markdown: {
	from: "github:tree-sitter-grammars/tree-sitter-markdown"
	repo: "tree-sitter-markdown"
	paths: ["tree-sitter-markdown/src", "tree-sitter-markdown-inline/src", "common"]
}

#grammar: matlab: {
	from: "github:acristoffers/tree-sitter-matlab"
	repo: "tree-sitter-matlab"
}

#grammar: meson: {
	from: "github:tree-sitter-grammars/tree-sitter-meson"
	repo: "tree-sitter-meson"
}

#grammar: move: {
	from: "github:tree-sitter-grammars/tree-sitter-move"
	repo: "tree-sitter-move"
}

#grammar: nasm: {
	from: "github:naclsn/tree-sitter-nasm"
	repo: "tree-sitter-nasm"
}

#grammar: nginx: {
	from: "github:opa-oz/tree-sitter-nginx"
	repo: "tree-sitter-nginx"
}

#grammar: nickel: {
	from: "github:nickel-lang/tree-sitter-nickel"
	repo: "tree-sitter-nickel"
}

#grammar: nim: {
	from: "github:alaviss/tree-sitter-nim"
	repo: "tree-sitter-nim"
}

#grammar: nix: {
	from: "github:nix-community/tree-sitter-nix"
	repo: "tree-sitter-nix"
}

#grammar: nqc: {
	from: "github:tree-sitter-grammars/tree-sitter-nqc"
	repo: "tree-sitter-nqc"
}

#grammar: objc: {
	from: "github:tree-sitter-grammars/tree-sitter-objc"
	repo: "tree-sitter-objc"
}

#grammar: objdump: {
	from: "github:ColinKennedy/tree-sitter-objdump"
	repo: "tree-sitter-objdump"
}

#grammar: ocaml: {
	from: "github:tree-sitter/tree-sitter-ocaml"
	repo: "tree-sitter-ocaml"
	paths: ["grammars/ocaml/src", "grammars/interface/src", "grammars/type/src", "common"]
}

#grammar: odin: {
	from: "github:tree-sitter-grammars/tree-sitter-odin"
	repo: "tree-sitter-odin"
}

#grammar: pascal: {
	from: "github:Isopod/tree-sitter-pascal"
	repo: "tree-sitter-pascal"
}

#grammar: pem: {
	from: "github:tree-sitter-grammars/tree-sitter-pem"
	repo: "tree-sitter-pem"
}

#grammar: perl: {
	from: "github:tree-sitter-perl/tree-sitter-perl"
	repo: "tree-sitter-perl"
}

#grammar: php: {
	from: "github:tree-sitter/tree-sitter-php"
	repo: "tree-sitter-php"
	paths: ["php/src", "php_only/src", "common"]
}

#grammar: pkl: {
	from: "github:apple/tree-sitter-pkl"
	repo: "tree-sitter-pkl"
}

#grammar: po: {
	from: "github:tree-sitter-grammars/tree-sitter-po"
	repo: "tree-sitter-po"
}

#grammar: "poe-filter": {
	from: "github:tree-sitter-grammars/tree-sitter-poe-filter"
	repo: "tree-sitter-poe-filter"
}

#grammar: pony: {
	from: "github:tree-sitter-grammars/tree-sitter-pony"
	repo: "tree-sitter-pony"
}

#grammar: powershell: {
	from: "github:airbus-cert/tree-sitter-powershell"
	repo: "tree-sitter-powershell"
}

#grammar: printf: {
	from: "github:tree-sitter-grammars/tree-sitter-printf"
	repo: "tree-sitter-printf"
}

#grammar: prisma: {
	from: "github:victorhqc/tree-sitter-prisma"
	repo: "tree-sitter-prisma"
}

#grammar: properties: {
	from: "github:tree-sitter-grammars/tree-sitter-properties"
	repo: "tree-sitter-properties"
}

#grammar: proto: {
	from: "github:mitchellh/tree-sitter-proto"
	repo: "tree-sitter-proto"
}

#grammar: puppet: {
	from: "github:tree-sitter-grammars/tree-sitter-puppet"
	repo: "tree-sitter-puppet"
}

#grammar: purescript: {
	from: "github:postsolar/tree-sitter-purescript"
	repo: "tree-sitter-purescript"
}

#grammar: pymanifest: {
	from: "github:tree-sitter-grammars/tree-sitter-pymanifest"
	repo: "tree-sitter-pymanifest"
}

#grammar: python: {
	from: "github:tree-sitter/tree-sitter-python"
	repo: "tree-sitter-python"
	astdump: {
		root_type: "module"
		first_child_type: "assignment"
	}
}

#grammar: ql: {
	from: "github:tree-sitter/tree-sitter-ql"
	repo: "tree-sitter-ql"
}

#grammar: "ql-dbscheme": {
	from: "github:tree-sitter/tree-sitter-ql-dbscheme"
	repo: "tree-sitter-ql-dbscheme"
}

#grammar: qmldir: {
	from: "github:tree-sitter-grammars/tree-sitter-qmldir"
	repo: "tree-sitter-qmldir"
}

#grammar: query: {
	from: "github:tree-sitter-grammars/tree-sitter-query"
	repo: "tree-sitter-query"
}

#grammar: racket: {
	from: "github:6cdh/tree-sitter-racket"
	repo: "tree-sitter-racket"
}

#grammar: rasi: {
	from: "github:Fymyte/tree-sitter-rasi"
	repo: "tree-sitter-rasi"
}

#grammar: re2c: {
	from: "github:tree-sitter-grammars/tree-sitter-re2c"
	repo: "tree-sitter-re2c"
}

#grammar: readline: {
	from: "github:tree-sitter-grammars/tree-sitter-readline"
	repo: "tree-sitter-readline"
}

#grammar: regex: {
	from: "github:tree-sitter/tree-sitter-regex"
	repo: "tree-sitter-regex"
}

#grammar: rego: {
	from: "github:FallenAngel97/tree-sitter-rego"
	repo: "tree-sitter-rego"
}

#grammar: requirements: {
	from: "github:tree-sitter-grammars/tree-sitter-requirements"
	repo: "tree-sitter-requirements"
}

#grammar: rescript: {
	from: "github:rescript-lang/tree-sitter-rescript"
	repo: "tree-sitter-rescript"
}

#grammar: roc: {
	from: "github:faldor20/tree-sitter-roc"
	repo: "tree-sitter-roc"
}

#grammar: ron: {
	from: "github:tree-sitter-grammars/tree-sitter-ron"
	repo: "tree-sitter-ron"
}

#grammar: ruby: {
	from: "github:tree-sitter/tree-sitter-ruby"
	repo: "tree-sitter-ruby"
}

#grammar: rust: {
	from: "github:tree-sitter/tree-sitter-rust"
	repo: "tree-sitter-rust"
}

#grammar: scala: {
	from: "github:tree-sitter/tree-sitter-scala"
	repo: "tree-sitter-scala"
}

#grammar: scheme: {
	from: "github:6cdh/tree-sitter-scheme"
	repo: "tree-sitter-scheme"
}

#grammar: scss: {
	from: "github:tree-sitter-grammars/tree-sitter-scss"
	repo: "tree-sitter-scss"
}

#grammar: slang: {
	from: "github:tree-sitter-grammars/tree-sitter-slang"
	repo: "tree-sitter-slang"
}

#grammar: smali: {
	from: "github:tree-sitter-grammars/tree-sitter-smali"
	repo: "tree-sitter-smali"
}

#grammar: smithy: {
	from: "github:indoorvivants/tree-sitter-smithy"
	repo: "tree-sitter-smithy"
}

#grammar: sml: {
	from: "github:MatthewFluet/tree-sitter-sml"
	repo: "tree-sitter-sml"
}

#grammar: solidity: {
	from: "github:JoranHonig/tree-sitter-solidity"
	repo: "tree-sitter-solidity"
}

#grammar: sparql: {
	from: "github:GordianDziwis/tree-sitter-sparql"
	repo: "tree-sitter-sparql"
}

#grammar: sql: {
	from: "github:DerekStride/tree-sitter-sql"
	repo: "tree-sitter-sql"
}

#grammar: squirrel: {
	from: "github:tree-sitter-grammars/tree-sitter-squirrel"
	repo: "tree-sitter-squirrel"
}

#grammar: "ssh-config": {
	from: "github:tree-sitter-grammars/tree-sitter-ssh-config"
	repo: "tree-sitter-ssh-config"
}

#grammar: starlark: {
	from: "github:tree-sitter-grammars/tree-sitter-starlark"
	repo: "tree-sitter-starlark"
}

#grammar: surface: {
	from: "github:connorlay/tree-sitter-surface"
	repo: "tree-sitter-surface"
}

#grammar: svelte: {
	from: "github:tree-sitter-grammars/tree-sitter-svelte"
	repo: "tree-sitter-svelte"
}

#grammar: swift: {
	from: "github:alex-pinkus/tree-sitter-swift"
	repo: "tree-sitter-swift"
}

#grammar: tablegen: {
	from: "github:tree-sitter-grammars/tree-sitter-tablegen"
	repo: "tree-sitter-tablegen"
}

#grammar: tcl: {
	from: "github:tree-sitter-grammars/tree-sitter-tcl"
	repo: "tree-sitter-tcl"
}

#grammar: templ: {
	from: "github:vrischmann/tree-sitter-templ"
	repo: "tree-sitter-templ"
}

#grammar: test: {
	from: "github:tree-sitter-grammars/tree-sitter-test"
	repo: "tree-sitter-test"
}

#grammar: thrift: {
	from: "github:tree-sitter-grammars/tree-sitter-thrift"
	repo: "tree-sitter-thrift"
}

#grammar: tmux: {
	from: "github:Freed-Wu/tree-sitter-tmux"
	repo: "tree-sitter-tmux"
}

#grammar: toml: {
	from: "github:tree-sitter/tree-sitter-toml"
	repo: "tree-sitter-toml"
}

#grammar: turtle: {
	from: "github:GordianDziwis/tree-sitter-turtle"
	repo: "tree-sitter-turtle"
}

#grammar: twig: {
	from: "github:gbprod/tree-sitter-twig"
	repo: "tree-sitter-twig"
}

#grammar: typescript: {
	from: "github:tree-sitter/tree-sitter-typescript"
	repo: "tree-sitter-typescript"
	paths: ["typescript/src", "tsx/src", "common"]
}

#grammar: typst: {
	from: "github:uben0/tree-sitter-typst"
	repo: "tree-sitter-typst"
}

#grammar: udev: {
	from: "github:tree-sitter-grammars/tree-sitter-udev"
	repo: "tree-sitter-udev"
}

#grammar: ungrammar: {
	from: "github:tree-sitter-grammars/tree-sitter-ungrammar"
	repo: "tree-sitter-ungrammar"
}

#grammar: uxntal: {
	from: "github:tree-sitter-grammars/tree-sitter-uxntal"
	repo: "tree-sitter-uxntal"
}

#grammar: verilog: {
	from: "github:tree-sitter/tree-sitter-verilog"
	repo: "tree-sitter-verilog"
}

#grammar: vhdl: {
	from: "github:alemuller/tree-sitter-vhdl"
	repo: "tree-sitter-vhdl"
}

#grammar: vim: {
	from: "github:tree-sitter-grammars/tree-sitter-vim"
	repo: "tree-sitter-vim"
}

#grammar: vimdoc: {
	from: "github:neovim/tree-sitter-vimdoc"
	repo: "tree-sitter-vimdoc"
}

#grammar: vue: {
	from: "github:tree-sitter-grammars/tree-sitter-vue"
	repo: "tree-sitter-vue"
}

#grammar: wasm: {
	from: "github:wasm-lsp/tree-sitter-wasm"
	repo: "tree-sitter-wasm"
	paths: ["wast/src", "wat/src"]
}

#grammar: wgsl: {
	from: "github:szebniok/tree-sitter-wgsl"
	repo: "tree-sitter-wgsl"
}

#grammar: "wgsl-bevy": {
	from: "github:tree-sitter-grammars/tree-sitter-wgsl-bevy"
	repo: "tree-sitter-wgsl-bevy"
}

#grammar: wit: {
	from: "github:bytecodealliance/tree-sitter-wit"
	repo: "tree-sitter-wit"
}

#grammar: xcompose: {
	from: "github:tree-sitter-grammars/tree-sitter-xcompose"
	repo: "tree-sitter-xcompose"
}

#grammar: xml: {
	from: "github:tree-sitter-grammars/tree-sitter-xml"
	repo: "tree-sitter-xml"
	paths: ["xml/src", "dtd/src", "common"]
}

#grammar: yaml: {
	from: "github:tree-sitter-grammars/tree-sitter-yaml"
	repo: "tree-sitter-yaml"
	paths: ["src", "schema/json/src", "schema/legacy/src", "schema/core/src"]
}

#grammar: yuck: {
	from: "github:tree-sitter-grammars/tree-sitter-yuck"
	repo: "tree-sitter-yuck"
}

#grammar: zig: {
	from: "github:tree-sitter-grammars/tree-sitter-zig"
	repo: "tree-sitter-zig"
}

#grammar: [string]: {
	from:    string
	version: string | *"HEAD"
	paths:   [...string] | *["src"]
	// Dest basename under third-party/ (usually tree-sitter-<key>).
	repo: string
	astdump?: {
		root_type:         string
		first_child_type?: string
	}
}

// Core C library. Placed next to the grammars; emcc reads that tree.
#tree_sitter: {
	from:    "github:tree-sitter/tree-sitter"
	version: "HEAD"
	repo:    "tree-sitter"
	paths: ["lib/src", "lib/include"]
}

workspaced: {
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
		tree_sitter: {
			from: "core:place"
			config: {
				ignore_missing: true
				items: {
					for p in #tree_sitter.paths {
						"third-party/\(#tree_sitter.repo)/\(p)": "tree_sitter:\(p)"
					}
				}
			}
		}
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
}

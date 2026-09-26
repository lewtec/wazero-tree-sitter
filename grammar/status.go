package grammar

import "sync"

var (
	liveParseOnce sync.Once
	liveParseOK   bool
)

// LiveParseReady reports whether lang can complete a parse.
// The probe runs at most once. Later calls ignore lang and return the cached result.
func LiveParseReady(lang *Language) bool {
	liveParseOnce.Do(func() {
		liveParseOK = probeLiveParse(lang)
	})
	return liveParseOK
}

func probeLiveParse(lang *Language) bool {
	if lang == nil {
		return false
	}
	p := NewParser()
	defer p.Delete()
	if !p.SetLanguage(lang) {
		return false
	}
	tree := p.ParseBytes([]byte("0"))
	defer tree.Delete()
	root := tree.RootNode()
	return root != nil && !root.IsNull()
}

package grammar

import "testing"

func TestRegisterGet(t *testing.T) {
	lang := &Language{wasm: []byte{0}}
	Register("json", lang)
	got, ok := Get("json")
	if !ok || got != lang {
		t.Fatalf("Get(json) = %v, %v", got, ok)
	}
	byExt, ok := GetByExtension("sample.JSON")
	if !ok || byExt != lang {
		t.Fatalf("GetByExtension = %v, %v", byExt, ok)
	}
	if SupportedLanguages() == "none" {
		t.Fatal("expected a registered language")
	}
}

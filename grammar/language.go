package grammar

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

var instanceSeq atomic.Uint64

func instanceName() string {
	return fmt.Sprintf("grammar-%d", instanceSeq.Add(1))
}

// Language is an immutable tree-sitter grammar compiled to wasm.
// Safe to share across parsers and goroutines.
type Language struct {
	wasm []byte

	mu       sync.Mutex
	rt       wazero.Runtime
	compiled wazero.CompiledModule
}

// NewLanguage wraps a wasm image built by codegen (emscripten standalone).
// The image is not instantiated until a Parser calls SetLanguage.
func NewLanguage(wasm []byte) *Language {
	return &Language{wasm: append([]byte(nil), wasm...)}
}

func (l *Language) module(ctx context.Context) (wazero.Runtime, wazero.CompiledModule, error) {
	if l == nil || len(l.wasm) == 0 {
		return nil, nil, fmt.Errorf("empty language wasm")
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.compiled != nil {
		return l.rt, l.compiled, nil
	}
	rt := wazero.NewRuntime(ctx)
	wasi_snapshot_preview1.MustInstantiate(ctx, rt)
	compiled, err := rt.CompileModule(ctx, l.wasm)
	if err != nil {
		_ = rt.Close(ctx)
		return nil, nil, err
	}
	l.rt = rt
	l.compiled = compiled
	return rt, compiled, nil
}

// guest is one wasm instance of a Language.
type guest struct {
	mod    api.Module
	malloc api.Function
	free   api.Function
	fns    map[string]api.Function
	nodeSz uint32
}

func instantiate(ctx context.Context, lang *Language) (*guest, error) {
	rt, compiled, err := lang.module(ctx)
	if err != nil {
		return nil, err
	}
	mod, err := rt.InstantiateModule(ctx, compiled, wazero.NewModuleConfig().WithName(instanceName()))
	if err != nil {
		return nil, err
	}
	g := &guest{mod: mod, fns: map[string]api.Function{}}
	for _, name := range []string{
		"malloc", "_malloc",
		"free", "_free",
	} {
		if fn := mod.ExportedFunction(name); fn != nil {
			switch name {
			case "malloc", "_malloc":
				if g.malloc == nil {
					g.malloc = fn
				}
			case "free", "_free":
				if g.free == nil {
					g.free = fn
				}
			}
		}
	}
	if g.malloc == nil || g.free == nil {
		_ = mod.Close(ctx)
		return nil, fmt.Errorf("wasm missing malloc/free exports")
	}
	sz, err := g.call1(ctx, "wts_node_sizeof")
	if err != nil {
		_ = mod.Close(ctx)
		return nil, err
	}
	if sz == 0 || sz > 256 {
		_ = mod.Close(ctx)
		return nil, fmt.Errorf("unexpected TSNode size %d", sz)
	}
	g.nodeSz = sz
	return g, nil
}

func (g *guest) close(ctx context.Context) {
	if g == nil || g.mod == nil {
		return
	}
	_ = g.mod.Close(ctx)
	g.mod = nil
}

func (g *guest) fn(name string) (api.Function, error) {
	if f, ok := g.fns[name]; ok {
		return f, nil
	}
	for _, n := range []string{name, "_" + name} {
		if f := g.mod.ExportedFunction(n); f != nil {
			g.fns[name] = f
			return f, nil
		}
	}
	return nil, fmt.Errorf("wasm missing export %s", name)
}

func (g *guest) call(ctx context.Context, name string, params ...uint64) ([]uint64, error) {
	f, err := g.fn(name)
	if err != nil {
		return nil, err
	}
	return f.Call(ctx, params...)
}

func (g *guest) call1(ctx context.Context, name string, params ...uint64) (uint32, error) {
	out, err := g.call(ctx, name, params...)
	if err != nil {
		return 0, err
	}
	if len(out) == 0 {
		return 0, nil
	}
	return uint32(out[0]), nil
}

func (g *guest) cstring(ptr uint32) string {
	if ptr == 0 || g.mod == nil || g.mod.Memory() == nil {
		return ""
	}
	mem := g.mod.Memory()
	var buf []byte
	for n := uint32(0); n < 1<<20; n++ {
		b, ok := mem.ReadByte(ptr + n)
		if !ok || b == 0 {
			break
		}
		buf = append(buf, b)
	}
	return string(buf)
}

func (g *guest) alloc(ctx context.Context, n uint32) (uint32, error) {
	if n == 0 {
		n = 1
	}
	return g.call1(ctx, "malloc", uint64(n))
}

func (g *guest) release(ctx context.Context, ptr uint32) {
	if ptr == 0 {
		return
	}
	_, _ = g.call(ctx, "free", uint64(ptr))
}

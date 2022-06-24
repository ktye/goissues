package main

import (
	"context"
	"fmt"
	"os"

	"github.com/tetratelabs/wazero"
)

func main() {
	file := os.Args[1]
	b, e := os.ReadFile(file)
	fatal(e)

	fmt.Printf("%s interpreted>\n", file)
	do(b, true)

	fmt.Printf("%s compiled>\n", file)
	do(b, true)

}
func do(wasm []byte, interpreted bool) {
	ctx := context.Background()

	var r wazero.Runtime
	if interpreted {
		r = wazero.NewRuntimeWithConfig(wazero.NewRuntimeConfigInterpreter().WithWasmCore2())
	} else {
		r = wazero.NewRuntimeWithConfig(wazero.NewRuntimeConfig().WithWasmCore2())
	}
	defer r.Close(ctx)

	m, e := r.InstantiateModuleFromBinary(ctx, wasm)
	fatal(e)

	res, e := m.ExportedFunction("kinit").Call(ctx)
	fatal(e)
	fmt.Printf("f() => %v\n", res)
}
func fatal(e error) {
	if e != nil {
		panic(e)
	}
}

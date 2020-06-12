package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"encoding/base64"

	"github.com/mathetake/gasm/wasm"
)

func main() {
	b, e := base64.StdEncoding.DecodeString("AGFzbQEAAAABBgFgAX8BfwMCAQAFAwEAAQcLAgNtZW0CAAFmAAAKKwEpAQF/IAAEQEEAIQEDQCABQQNLBEAgAQ8LIAFBAWoiASAASQ0ACwsgAQs=")
	check(e)

	m, e := wasm.DecodeModule(bytes.NewReader(b))
	check(e)

	vm, e := wasm.NewVM(m, nil)
	check(e)

	r, _, e := vm.ExecExportedFunction("f", 10)
	check(e)

	fmt.Println("result", r)
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func gen() {
	b, e := ioutil.ReadFile("min.wasm")
	check(e)
	fmt.Println(base64.StdEncoding.EncodeToString(b))
}

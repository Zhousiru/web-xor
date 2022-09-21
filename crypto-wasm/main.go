package main

import (
	"encoding/base64"
	"syscall/js"
)

func crypto(this js.Value, args []js.Value) interface{} {
	buffer := make([]byte, args[0].Length())
	js.CopyBytesToGo(buffer, args[0])

	bKey := []byte(args[1].String())

	keyLen := len(bKey)
	r := make([]byte, args[0].Length())

	for index, v := range buffer {
		r[index] = v ^ bKey[index%keyLen]
	}

	return base64.StdEncoding.EncodeToString(r)
}

func main() {
	done := make(chan struct{})
	js.Global().Get("wasm").Set("crypto", js.FuncOf(crypto))
	<-done
}

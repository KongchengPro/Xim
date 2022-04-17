//go:build js || wasm

package xim

import "syscall/js"

//goland:noinspection GoUnusedGlobalVariable
var (
	doc  = js.Global().Get("document")
	body = doc.Get("body")
)

func SetTitle(title string) {
	doc.Set("title", title)
}

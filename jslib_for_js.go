//go:build js || wasm

package xim

import "syscall/js"

type JsValue = js.Value

type JsFunc = js.Func

func JsFuncOf(fn func(this JsValue, args []JsValue) any) JsFunc {
	return js.FuncOf(fn)
}

func JsNilValue() JsValue {
	return js.ValueOf(nil)
}

//goland:noinspection GoUnusedGlobalVariable
var (
	doc  = js.Global().Get("document")
	body = doc.Get("body")
)

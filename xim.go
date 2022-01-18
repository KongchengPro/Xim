package xim

import (
	. "gitee.com/kogic/xim/component"
	. "gitee.com/kogic/xim/style"
	"syscall/js"
)

var (
	doc        = js.Global().Get("document")
	body       = doc.Get("body")
	console    = js.Global().Get("console")
	consoleLog = func(args ...interface{}) { console.Call("log", args...) }
)

func Render(v View) {

	for _, component := range v.Components {
		elStr, contentStr, style := component.Generate()
		el := doc.Call("createElement", elStr)
		el.Set("innerHTML", contentStr)
		el.Call("setAttribute", "style", GenerateCSS(style))
		consoleLog(el.Call("getAttribute", "style"))
		body.Call("appendChild", el)
	}
}

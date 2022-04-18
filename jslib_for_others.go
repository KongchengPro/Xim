//go:build !(js || wasm)

package xim

type JsValue struct {
}

func (v *JsValue) Get(key string) JsValue {
	return JsValue{}
}

func (v *JsValue) Set(key string, value any) {

}

func (v JsValue) Call(args ...interface{}) JsValue {
	return JsValue{}
}

type JsFunc struct {
}

func JsNilValue() JsValue {
	return JsValue{}
}

func JsFuncOf(fn func(this JsValue, args []JsValue) any) JsFunc {
	return JsFunc{}
}

//goland:noinspection GoUnusedGlobalVariable
var (
	doc  = JsValue{}
	body = JsValue{}
)

//go:build !(js || wasm)

package xim

type JsValue struct {
}

//goland:noinspection GoUnusedParameter
func (v JsValue) Get(key string) JsValue {
	return JsValue{}
}

//goland:noinspection GoUnusedParameter
func (v JsValue) Index(index int) JsValue {
	return JsValue{}
}

//goland:noinspection GoUnusedParameter
func (v JsValue) Set(key string, value any) {

}

//goland:noinspection GoUnusedParameter
func (v JsValue) Call(args ...interface{}) JsValue {
	return JsValue{}
}

func (v JsValue) String() string {
	return ""
}

type JsFunc struct {
}

func JsNilValue() JsValue {
	return JsValue{}
}

//goland:noinspection GoUnusedParameter
func JsFuncOf(fn func(this JsValue, args []JsValue) any) JsFunc {
	return JsFunc{}
}

//goland:noinspection GoUnusedGlobalVariable
var (
	global = JsValue{}
	doc    = JsValue{}
	body   = JsValue{}
	window = JsValue{}
)

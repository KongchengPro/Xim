package xim

import (
	"github.com/kongchengpro/xim/api/dom"
	"github.com/kongchengpro/xim/internal/constants"
	. "github.com/kongchengpro/xim/types/component"
	. "github.com/kongchengpro/xim/types/style"
	"github.com/sirupsen/logrus"
	"syscall/js"
)

//goland:noinspection GoUnusedGlobalVariable
var (
	doc  = js.Global().Get("document")
	body = doc.Get("body")
)

func Init(c Component, parent Component) {
	// 找到父元素
	var parentElement js.Value
	if parent.GetID() == constants.RootComponentID {
		parentElement = doc.Get("body")
	} else {
		parentElement = doc.Call("getElementById", parent.GetID())
	}
	if parentElement.Equal(js.ValueOf(nil)) {
		logrus.Errorf("找不到父组件：%T%+v", parent, parent)
		return
	}
	// 创建元素
	elementType := c.GetElementType()
	element := doc.Call("createElement", elementType)
	element.Set("id", c.GetID())
	callbackMap := c.GetCallbackMap()
	if callbackMap != nil {
		for callbackName, callbackFunc := range callbackMap {
			element.Call("addEventListener", callbackName, js.FuncOf(func(this js.Value, args []js.Value) interface{} {
				callbackFunc(this, args...)
				return nil
			}))
		}
	}
	// 在父元素中添加
	parentElement.Call("appendChild", element)
	// 刷新一次组件以显示内容
	Refresh(c)
	// 在组件映射中添加
	dom.ComponentMap[parent.GetName()+c.GetName()] = c
	// 递归渲染子组件
	if len(c.GetChildComponents()) != 0 {
		for _, childComponent := range c.GetChildComponents() {
			Init(childComponent, c)
		}
	}
}

// Refresh 刷新组件显示的内容（不包括回调）
func Refresh(components ...Component) {
	for _, c := range components {
		element := doc.Call("getElementById", c.GetID())
		// 获取组件相关信息
		innerHTML := c.GetInnerHTML()
		style := c.GetStyle()
		if innerHTML != "" {
			element.Set("innerHTML", innerHTML)
		}
		if style != nil {
			element.Call("setAttribute", "style", GenerateCSS(style))
		}
	}
}

func SetTitle(title string) {
	doc.Set("title", title)
}

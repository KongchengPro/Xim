package xim

import (
	"gitee.com/kogic/xim/internal/constants"
	. "gitee.com/kogic/xim/types"
	"github.com/sirupsen/logrus"
	"syscall/js"
)

//goland:noinspection GoUnusedGlobalVariable
var (
	doc  = js.Global().Get("document")
	body = doc.Get("body")
)

// Render 递归渲染组件
func Render(c Component, parent Component) {
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
	// 获取组件相关信息
	elementType := c.GetElementType()
	innerHTML := c.GetInnerHTML()
	style := c.GetStyle()
	// 创建元素
	element := doc.Call("createElement", elementType)
	element.Set("id", c.GetID())
	if innerHTML != "" {
		element.Set("innerHTML", innerHTML)
	}
	if style != nil {
		element.Call("setAttribute", "style", GenerateCSS(style))
	}
	// 在父元素中添加
	parentElement.Call("appendChild", element)
	if len(c.GetChildComponents()) != 0 {
		for _, childComponent := range c.GetChildComponents() {
			Render(childComponent, c)
		}
	}
}

func SetTitle(title string) {
	doc.Set("title", title)
}

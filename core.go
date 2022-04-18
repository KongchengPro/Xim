package xim

import (
	"fmt"
	"github.com/LouisCheng-CN/xim/components"
	"github.com/LouisCheng-CN/xim/types"
)

func GenerateRawComponentTree(ctx *types.Context, rc *types.RawComponent) error {
	if rc.Children == nil {
		rc.Children = make([]*types.RawComponent, 0)
	}
	for _, childCtx := range ctx.Children {
		childComp := childCtx.Component
		switch typedChildComp := childComp.(type) {
		case components.Panel:
			childRc := &types.RawComponent{
				LabelName: "div",
				Attribute: map[string]string{
					"style": "background:" + typedChildComp.Color + ";",
				},
			}
			err := GenerateRawComponentTree(childCtx, childRc)
			if err != nil {
				return err
			}
			rc.Children = append(rc.Children, childRc)
		case components.Text:
			rc.Children = append(rc.Children, &types.RawComponent{
				LabelName: "p",
				Attribute: nil,
				Content:   typedChildComp.Content,
			})
		default:
			err := GenerateRawComponentTree(childCtx, rc)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func SetContent(content types.Component) {
	rawComponentTree := types.RawComponent{
		LabelName: "div",
		Attribute: nil,
		Content:   "",
	}
	rootCtx := &types.Context{}
	content.Compose(rootCtx)
	err := GenerateRawComponentTree(rootCtx, &rawComponentTree)
	if err != nil {
		fmt.Println(err)
	}
}

//
//func AddChild(child Composable, parent Composable) {
//	comp := child.Compose()
//	// 找到父元素
//	var parentElement js.Value
//	if _, ok := parent.(*components.Root); ok {
//		parentElement = doc.Get("body")
//	} else {
//		parentElement = doc.Call("getElementById", parent.Compose().Id())
//	}
//	if parentElement.Equal(js.ValueOf(nil)) {
//		fmt.Println("找不到父组件：%T%+v", parent, parent)
//		return
//	}
//	// 创建元素
//	var elemType string
//	if comp.IsInline {
//		elemType = "div"
//	} else {
//		elemType = "span"
//	}
//	elem := doc.Call("createElement", elemType)
//	elem.Set("id", comp.Id())
//	callbackMap := comp.CallbackMap
//	if callbackMap != nil {
//		for callbackName, callbackFunc := range callbackMap {
//			elem.Call("addEventListener", callbackName, js.FuncOf(func(this js.Value, args []js.Value) interface{} {
//				callbackFunc(this, args...)
//				return nil
//			}))
//		}
//	}
//	// 在父元素中添加
//	parentElement.Call("appendChild", elem)
//	// 刷新一次组件以显示内容
//	Refresh(comp)
//	// 递归渲染子组件
//	if len(comp.Children) != 0 {
//		for _, child := range comp.Children {
//			AddChild(child, comp)
//		}
//	}
//}
//
//// Refresh 刷新组件的内容和属性
//func Refresh(comp *Composition) {
//	element := doc.Call("getElementById", comp.Id())
//	html := comp.Html
//	attrs := comp.Attrs
//	if html != "" {
//		element.Set("innerHTML", html)
//	}
//	for key, value := range attrs {
//		element.Call("setAttribute", key, value)
//	}
//	if comp.Children != nil {
//		for _, child := range comp.Children {
//			Refresh(child.Compose())
//		}
//	}
//}

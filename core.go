package xim

import (
	"encoding/json"
	"fmt"
	"github.com/LouisCheng-CN/xim/components"
	"github.com/LouisCheng-CN/xim/types"
)

var IsDebug bool

func SetTitle(title string) {
	doc.Set("title", title)
}

func SetContent(content types.Component) {
	rawComponentTree := &types.RawComponent{}
	rootCtx := &types.Context{}
	content.Compose(rootCtx)
	if IsDebug {
		bs, _ := json.MarshalIndent(rootCtx, "", "  ")
		fmt.Printf("context: %+v\n", string(bs))
	}
	err := generateRawComponentTree(rootCtx, rawComponentTree)
	if err != nil {
		fmt.Println("generateRawComponentTree error:", err)
		return
	}
	if IsDebug {
		bs, _ := json.MarshalIndent(rawComponentTree, "", "  ")
		fmt.Println("rawComponentTree:", string(bs))
	}
	renderRawComponentTree(rawComponentTree, JsNilValue(), true)
	go func() {
		for {
			select {
			case subscriber := <-types.RefreshChannel:
				comp := rootCtx.Search(subscriber)
				newCtx := &types.Context{}
				comp.Compose(newCtx)
				newRawComponentTree := &types.RawComponent{}
				err := generateRawComponentTree(newCtx, newRawComponentTree)
				if err != nil {
					fmt.Println("generateRawComponentTree error [new]:", err)
					return
				}
				if IsDebug {
					bs, _ := json.MarshalIndent(newRawComponentTree, "", "  ")
					fmt.Println("newRawComponentTree [new]:", string(bs))
				}
				//renderRawComponentTree(rawComponentTree, JsNilValue(), true)
			default:
			}
		}
	}()
}

func generateRawComponentTree(ctx *types.Context, rc *types.RawComponent) error {
	if len(ctx.Children) != 0 && rc.Children == nil {
		rc.Children = make([]*types.RawComponent, 0)
	}
	for _, childCtx := range ctx.Children {
		var childRc *types.RawComponent
		childComp := childCtx.Component
		switch typedChildComp := childComp.(type) {
		case *components.Panel:
			childRc = &types.RawComponent{
				LabelName: "div",
				Attributes: map[string]string{
					"style": "background:" + typedChildComp.Color + ";",
				},
			}
			err := generateRawComponentTree(childCtx, childRc)
			if err != nil {
				return err
			}

		case *components.Text:
			childRc = &types.RawComponent{
				LabelName:  "p",
				Attributes: nil,
				Content:    typedChildComp.Content,
			}
		case *components.Button:
			childRc = &types.RawComponent{
				LabelName:  "button",
				Attributes: nil,
				Content:    typedChildComp.Content,
			}
		default:
			err := generateRawComponentTree(childCtx, rc)
			if err != nil {
				return err
			}
		}
		childRc.Id = childComp.Id()
		childRc.EventListeners = childCtx.EventListeners
		rc.Children = append(rc.Children, childRc)
	}
	return nil
}

func renderRawComponentTree(rc *types.RawComponent, parentElem JsValue, isRoot bool) {
	if isRoot {
		parentElem = doc.Get("body")
		for _, childRc := range rc.Children {
			renderRawComponentTree(childRc, parentElem, false)
		}
		return
	}
	if IsDebug {
		//fmt.Printf("renderRawComponentTree_createElement: %+v\n", rc)
	}
	elem := doc.Call("createElement", rc.LabelName)
	//elem.Set("id", comp.Id())
	elem.Set("innerHTML", rc.Content)
	if rc.EventListeners != nil {
		for event, handler := range rc.EventListeners {
			elem.Call("addEventListener", event, JsFuncOf(func(this JsValue, args []JsValue) interface{} {
				handler()
				return nil
			}))
		}
	}
	// 在父元素中添加
	parentElem.Call("appendChild", elem)
	// 递归渲染子组件
	if len(rc.Children) != 0 {
		for _, child := range rc.Children {
			renderRawComponentTree(child, elem, false)
		}
	}
}

//Refresh 刷新组件的内容和属性
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

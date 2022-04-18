package xim

import (
	"encoding/json"
	"fmt"
	"github.com/LouisCheng-CN/xim/components"
	"github.com/LouisCheng-CN/xim/internal/utils"
	"github.com/LouisCheng-CN/xim/types"
)

var IsDebug bool

func SetTitle(title string) {
	doc.Set("title", title)
}

func SetContent(content types.Component) {
	rootCtx := &types.Context{}
	content.Compose(rootCtx)
	if IsDebug {
		bs, _ := json.MarshalIndent(rootCtx, "", "  ")
		fmt.Printf("context: %+v\n", string(bs))
	}
	rawComponentTree, err := generateRawComponentTree(rootCtx)
	if err != nil {
		fmt.Println("generateRawComponentTree error:", err)
		return
	}
	if IsDebug {
		bs, _ := json.MarshalIndent(rawComponentTree, "", "  ")
		fmt.Println("rawComponentTree:", string(bs))
	}
	renderRawComponentTree(rawComponentTree, JsNilValue(), true, false, JsNilValue())
	go func() {
		for {
			select {
			case subscriber := <-types.RefreshChannel:
				indexList, ctx := rootCtx.Find(subscriber)
				comp := ctx.Component
				newCtx := &types.Context{}
				if IsDebug {
					utils.PrintlnStruct("comp [new]:", comp)
				}
				comp.Compose(newCtx)
				if IsDebug {
					bs, _ := json.MarshalIndent(newCtx, "", "  ")
					fmt.Println("newCtx:", string(bs))
				}
				newRawComponentTree, err := generateRawComponentTree(newCtx)
				if err != nil {
					fmt.Println("generateRawComponentTree error [new]:", err)
					return
				}
				if IsDebug {
					bs, _ := json.MarshalIndent(newRawComponentTree, "", "  ")
					fmt.Println("newRawComponentTree [new]:", string(bs))
				}
				contentRootDiv := body.Get("children").Index(3)
				target := contentRootDiv
				targetParent := contentRootDiv
				for i, index := range indexList {
					target = target.Get("children").Index(index)
					if i != len(indexList)-1 {
						targetParent = targetParent.Get("children").Index(index)
					}
				}
				//targetParent.Call("removeChild", target)
				renderRawComponentTree(newRawComponentTree, targetParent, false, true, target)
			}
		}
	}()
}

func generateRawComponentTree(ctx *types.Context) (*types.RawComponent, error) {
	if ctx.Children == nil {
		ctx.Children = make([]*types.Context, 0)
	}
	comp := ctx.Component
	rc := &types.RawComponent{
		Children: make([]*types.RawComponent, 0),
	}
	if comp == nil {
		for _, childCtx := range ctx.Children {
			childRc, err := generateRawComponentTree(childCtx)
			if err != nil {
				return nil, err
			}
			rc.Children = append(rc.Children, childRc)
		}
		return rc, nil
	}
	switch typedChildComp := comp.(type) {
	case *components.Panel:
		rc = &types.RawComponent{
			LabelName: "div",
			Attributes: map[string]string{
				"style": "background:" + typedChildComp.Color + ";",
			},
		}
		for _, childCtx := range ctx.Children {
			childRc, err := generateRawComponentTree(childCtx)
			if err != nil {
				return nil, err
			}
			rc.Children = append(rc.Children, childRc)
		}
	case *components.Text:
		rc = &types.RawComponent{
			LabelName:  "p",
			Attributes: nil,
			Content:    typedChildComp.Content,
		}
	case *components.Button:
		rc = &types.RawComponent{
			LabelName:  "button",
			Attributes: nil,
			Content:    typedChildComp.Content,
		}
	default:
		for _, childCtx := range ctx.Children {
			childRc, err := generateRawComponentTree(childCtx)
			if err != nil {
				return nil, err
			}
			rc.Children = append(rc.Children, childRc)
		}
	}
	rc.Id = comp.Id()
	rc.EventListeners = ctx.EventListeners
	return rc, nil
}

func renderRawComponentTree(rc *types.RawComponent, parentElem JsValue, isRoot bool, isReplace bool, oldElem JsValue) {
	if isRoot {
		parentElem = doc.Get("body")
		elem := doc.Call("createElement", "div")
		parentElem.Call("appendChild", elem)
		for _, childRc := range rc.Children {
			renderRawComponentTree(childRc, elem, false, false, JsNilValue())
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
	if isReplace {
		parentElem.Call("replaceChild", elem, oldElem)
	} else {
		parentElem.Call("appendChild", elem)
	}
	// 递归渲染子组件
	if len(rc.Children) != 0 {
		for _, child := range rc.Children {
			renderRawComponentTree(child, elem, false, false, JsNilValue())
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

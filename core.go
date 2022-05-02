package xim

import (
	"encoding/json"
	"fmt"
	"github.com/Re-Ch-Love/xim/internal/utils"
	"github.com/Re-Ch-Love/xim/types"
)

var IsDebug bool

func SetTitle(title string) {
	doc.Set("title", title)
}

func getViewRoot() JsValue {
	return body.Get("lastElementChild")
}

// SetContent 设置内容
// 会清除已有的内容
func SetContent(content types.Component) {
	getViewRoot().Call("remove")
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
	initRenderRawComponentTree(rawComponentTree)
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
				contentRootDiv := getViewRoot()
				target := contentRootDiv
				targetParent := contentRootDiv
				for i, index := range indexList {
					target = target.Get("children").Index(index)
					if i != len(indexList)-1 {
						targetParent = targetParent.Get("children").Index(index)
					}
				}
				replaceRenderRawComponentTree(newRawComponentTree, targetParent, target)
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
	render, ok := comp.(types.Render)
	if ok {
		internalRc, isRecursive := render.Render()
		if isRecursive {
			for _, childCtx := range ctx.Children {
				childRc, err := generateRawComponentTree(childCtx)
				if err != nil {
					return nil, err
				}
				internalRc.Children = append(internalRc.Children, childRc)
			}
		}
		rc = internalRc
	} else {
		for _, childCtx := range ctx.Children {
			childRc, err := generateRawComponentTree(childCtx)
			if err != nil {
				return nil, err
			}
			rc.Children = append(rc.Children, childRc)
		}
	}
	//switch typedChildComp := comp.(type) {
	//case *components.Panel:
	//	rc = &types.RawComponent{
	//		LabelName: "div",
	//		Attributes: map[string]string{
	//			"style": "background:" + typedChildComp.Color + ";",
	//		},
	//	}
	//	for _, childCtx := range ctx.Children {
	//		childRc, err := generateRawComponentTree(childCtx)
	//		if err != nil {
	//			return nil, err
	//		}
	//		rc.Children = append(rc.Children, childRc)
	//	}
	//case *components.Text:
	//	rc = &types.RawComponent{
	//		LabelName:  "p",
	//		Attributes: nil,
	//		Content:    typedChildComp.Content.Calculate(),
	//	}
	//case *components.Button:
	//	rc = &types.RawComponent{
	//		LabelName:  "button",
	//		Attributes: nil,
	//		Content:    typedChildComp.Content,
	//	}
	//default:
	//	for _, childCtx := range ctx.Children {
	//		childRc, err := generateRawComponentTree(childCtx)
	//		if err != nil {
	//			return nil, err
	//		}
	//		rc.Children = append(rc.Children, childRc)
	//	}
	//}
	rc.Id = comp.Id()
	rc.EventListeners = ctx.EventListeners
	return rc, nil
}

func replaceRenderRawComponentTree(rc *types.RawComponent, parentElem JsValue, oldElem JsValue) {
	elem := createElementFromRawComponent(rc)
	parentElem.Call("replaceChild", elem, oldElem)
	if len(rc.Children) != 0 {
		for _, child := range rc.Children {
			renderRawComponentTree(child, elem)
		}
	}
}

func renderRawComponentTree(rc *types.RawComponent, parentElem JsValue) {
	elem := createElementFromRawComponent(rc)
	for key, value := range rc.Attributes {
		elem.Set(key, value)
	}
	// 在父元素中添加
	parentElem.Call("appendChild", elem)
	// 递归渲染子组件
	if len(rc.Children) != 0 {
		for _, child := range rc.Children {
			renderRawComponentTree(child, elem)
		}
	}
}

func initRenderRawComponentTree(rc *types.RawComponent) {
	parentElem := doc.Get("body")
	elem := doc.Call("createElement", "div")
	parentElem.Call("appendChild", elem)
	for _, childRc := range rc.Children {
		renderRawComponentTree(childRc, elem)
	}
}

func createElementFromRawComponent(rc *types.RawComponent) JsValue {
	elem := doc.Call("createElement", rc.LabelName)
	elem.Set("innerHTML", rc.Content)
	if rc.EventListeners != nil {
		for event, handler := range rc.EventListeners {
			elem.Call("addEventListener", event, JsFuncOf(func(this JsValue, args []JsValue) interface{} {
				handler()
				return nil
			}))
		}
	}
	return elem
}

// Finish blocks the goroutine
func Finish() {
	select {}
}

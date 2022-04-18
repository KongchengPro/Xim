package xim

import (
	"encoding/json"
	"fmt"
	. "github.com/LouisCheng-CN/xim/components"
	"github.com/LouisCheng-CN/xim/types"
	"testing"
)

type HelloWorld struct {
	children []types.Component
	types.BaseComponent
}

func (h HelloWorld) Compose(ctx *types.Context) {
	ctx.Apply(h.children)
}

func TestGenerateRawComponentTree(t *testing.T) {
	var flag = false
	var HelloWorldComp = HelloWorld{
		children: []types.Component{
			&Panel{
				Color: "#fcfaed",
				Children: []types.Component{
					&Text{
						Content: "Hello $name!",
					},
					&Button{
						Content: "Click me!",
						OnClick: func() {
							flag = true
						},
					},
				},
			},
		},
	}
	var rawComp = &types.RawComponent{
		LabelName:  "",
		Attributes: nil,
		Content:    "",
	}
	ctx := &types.Context{}
	HelloWorldComp.Compose(ctx)
	err := generateRawComponentTree(ctx, rawComp)
	if err != nil {
		panic(err)
	}
	rawCompJsonBs, err := json.MarshalIndent(rawComp, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(rawCompJsonBs))
	var target = &types.RawComponent{
		LabelName:  "",
		Attributes: nil,
		Content:    "",
		Children: []*types.RawComponent{
			{
				LabelName: "div",
				Attributes: map[string]string{
					"style": "background:#fcfaed;",
				},
				Children: []*types.RawComponent{
					{
						LabelName:  "p",
						Attributes: nil,
						Content:    "Hello $name!",
					},
					{
						LabelName:  "button",
						Attributes: nil,
						Content:    "Click me!",
					},
				},
			},
		},
	}
	if err != nil {
		panic(err)
	}
	targetJsonBs, err := json.MarshalIndent(target, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(targetJsonBs))
	rawCompBtnHandler := rawComp.Children[0].Children[1].EventListeners["click"]
	rawCompBtnHandler()
	if flag && string(rawCompJsonBs) != string(targetJsonBs) {
		t.Fail()
	}
}

//ComposeString("Hello $Name", "name")
//
//type StrValue interface {
//	~string | ~func(f string, args ...string) string
//}
//
//var globalVar = make(map[string]types.Value)
//
//func ComposeString(f string, args ...string) func() string {
//	return func() string {
//		something := make([]interface{}, 0)
//		for _, arg := range args {
//			something = append(something, globalVar[arg])
//		}
//		sp := fmt.Sprintf(f, something...)
//		return sp
//	}
//}

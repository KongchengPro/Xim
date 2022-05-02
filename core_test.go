package xim

import (
	"encoding/json"
	"fmt"
	. "github.com/Re-Ch-Love/xim/components"
	"github.com/Re-Ch-Love/xim/types"
	"testing"
)

type HelloWorld struct {
	children []types.Component
	types.BaseComponent
}

func (h HelloWorld) Compose(ctx *types.Context) {
	ctx.AddChildren(h.children)
}

func TestGenerateRawComponentTree(t *testing.T) {
	var flag = false
	var HelloWorldComp = HelloWorld{
		children: []types.Component{
			&Panel{
				Color: "#fcfaed",
				Children: []types.Component{
					&Text{
						Content: types.NewStaticData[string]("Hello World"),
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
	ctx := &types.Context{}
	HelloWorldComp.Compose(ctx)
	rawComp, err := generateRawComponentTree(ctx)
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
//var globalVar = make(map[string]types.Data)
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

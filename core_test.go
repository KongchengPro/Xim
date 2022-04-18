package xim

import (
	. "github.com/LouisCheng-CN/xim/components"
	"github.com/LouisCheng-CN/xim/types"
	"reflect"
	"testing"
)

type HelloWorld struct {
	children []types.Component
	types.BaseComponent
}

func (h HelloWorld) Compose(ctx *types.Context) {
	ctx.Apply(h.children)
}

func TestRender(t *testing.T) {
	var HelloWorldComp = HelloWorld{
		children: []types.Component{
			Panel{
				Color: "#fcfaed",
				Children: []types.Component{
					Text{
						Content: "Hello $name!",
					},
				},
			},
		},
	}
	var rawComp = &types.RawComponent{
		LabelName: "",
		Attribute: nil,
		Content:   "",
	}
	ctx := &types.Context{}
	HelloWorldComp.Compose(ctx)
	err := GenerateRawComponentTree(ctx, rawComp)
	if err != nil {
		panic(err)
	}
	//bs, _ := json.MarshalIndent(rawComp, "", "  ")
	//fmt.Println(string(bs))
	var target = &types.RawComponent{
		LabelName: "",
		Attribute: nil,
		Content:   "",
		Children: []*types.RawComponent{
			{
				LabelName: "div",
				Attribute: map[string]string{
					"style": "background:#fcfaed;",
				},
				Children: []*types.RawComponent{
					{
						LabelName: "p",
						Attribute: nil,
						Content:   "Hello $name!",
					},
				},
			},
		},
	}
	//fmt.Println("generate:")
	//fmt.Println(rawComp)
	//fmt.Println("target:")
	//fmt.Println(target)
	if !reflect.DeepEqual(rawComp, target) {
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

package hello_world

import (
	"fmt"
	"github.com/Re-Ch-Love/xim"
	. "github.com/Re-Ch-Love/xim/components"
	. "github.com/Re-Ch-Love/xim/types"
)

type states struct {
	name MutableState[string]
}

var storage = &Storage[states]{
	States: &states{
		name: MutableStateOf("World"),
	},
	Mutations: map[string]func(*states, ...any){
		"setName": func(s *states, args ...any) {
			s.name.SetValue(args[0].(string))
		},
	},
}

var HelloWorld = Panel{
	Children: []Component{
		Panel{
			Color: "#F2F2F2",
			Children: []Component{
				Text{
					Initializer: func(text *Text) {
						text.Content = NewDynamicData[string](func() string {
							return "Hello " + storage.States.name.Value(text.Id()) + "!"
						})
					},
				}.Create(),
				Button{
					Content: "Click me!",
					OnClick: func() {
						storage.Commit("setName", "Xim")
					},
				}.Create(),
				Button{
					Content: "Go To Counter",
					OnClick: func() {
						err := xim.JumpTo("counter")
						if err != nil {
							fmt.Println(err)
							return
						}
					},
				}.Create(),
			},
		}.Create(),
	},
}.Create()

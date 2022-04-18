package components

import (
	"fmt"
	. "github.com/LouisCheng-CN/xim/components"
	. "github.com/LouisCheng-CN/xim/types"
)

type helloWorld struct {
	children []Component
	BaseComponent
}

func (h helloWorld) Compose(ctx *Context) {
	ctx.Apply(h.children)
}

type states struct {
	name MutableState[string]
}

var storage = &Storage[states]{
	States: &states{
		name: MutableStateOf("World"),
	},
	Mutations: map[string]func(*states, ...Value){
		"setName": func(s *states, args ...Value) {
			s.name.SetValue(args[0].(string))
		},
	},
}

var HelloWorld = helloWorld{
	children: []Component{
		&Panel{
			Color: "#fcfaed",
			Children: []Component{
				New[Text](func(t *Text) {
					t.DynamicContent = func() string {
						return "Hello " + storage.States.name.Value(t.Id()) + "!"
					}
				}),
				&Button{
					Content: "Click me!",
					OnClick: func() {
						fmt.Println("Clicked!")
						storage.Commit("setName", "Xim")
					},
				},
			},
		},
	},
}

/*
Component {
	children: {
		"MainText": Text {
		}
		"Button": Button {
		}
	},
	OnCreate: func(ctx *Context, t *children) {

	},
	Compose: func(ctx *Context, t *children) {
		mainText := t["MainText"]
		button := t["Button"]
		ctx.repeat(mainText, expr)
		ctx.create(button)
	},
}

Component Page {
	children: {
		"Title": Title {
		}
		"Content": Content {
		}
	},
	Compose: func(ctx *Context, t *children) {
		ctx.create(t["Title"])
		ctx.create(t["Content"])
	},
}

xim.SetContent(Page())
*/

/*
let(
	Add2,
	fn(int->int, arg => body)
)
let-fn(
	Add2,
	int -> [int -> int],
	arg => fn(arg => Add(arg, 2))
)

A(a, B(b, c))
A-B(a, b, c)
*/

/*
fn(..., .... ,....)
|--------------------------------------|
|                                      |
|                                      |
|           |^^^^^^^^^^^^^^|           |
|           | fn(..., ...) |           |
|           |______________|           |
|             ____fn____               |
|                                      |
|--------------------------------------|
					|
            ------------------
*/

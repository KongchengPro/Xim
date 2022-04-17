//go:build js || wasm

package main

import (
	. "github.com/LouisCheng-CN/xim"
	. "github.com/LouisCheng-CN/xim/components"
	. "github.com/LouisCheng-CN/xim/types"
)

type HelloWorld struct {
	children []Component
	BaseComponent
}

func (h HelloWorld) Compose(ctx *Context) {
	ctx.Apply(h.children)
}

var HelloWorldComp = HelloWorld{
	children: []Component{
		Panel{
			Color: "#fcfaed",
			Children: []Component{
				Text{
					Content: "Hello $name!",
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

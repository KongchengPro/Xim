package components

import "github.com/LouisCheng-CN/xim/types"

// Button 按钮
type Button struct {
	Content string
	OnClick func()
	types.BaseComponent
}

func (b Button) Compose(ctx *types.Context) {
	ctx.AddEventListener("click", b.OnClick)
}

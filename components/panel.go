package components

import (
	"github.com/LouisCheng-CN/xim/types"
)

//Panel 面板
type Panel struct {
	Color    string
	Children []types.Component
	types.BaseComponent
}

func (p *Panel) Compose(ctx *types.Context) {
	ctx.Component = p
	ctx.Apply(p.Children)
}

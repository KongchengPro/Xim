package components

import (
	"github.com/LouisCheng-CN/xim/types"
)

// Text 文本
type Text struct {
	Content        string
	DynamicContent func() string
	types.BaseComponent
}

func (t *Text) Compose(ctx *types.Context) {
	if t.DynamicContent != nil {
		t.Content = t.DynamicContent()
	}
	ctx.Component = t
}

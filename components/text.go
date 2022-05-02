package components

import (
	"github.com/Re-Ch-Love/xim/types"
)

// Text 文本
type Text struct {
	Content     types.Data[string]
	Initializer func(*Text)
	types.BaseComponent
}

func (t Text) Render() (*types.RawComponent, bool) {
	return &types.RawComponent{
		LabelName:  "p",
		Attributes: nil,
		Content:    t.Content.Calculate(),
	}, false
}

func (t Text) Create() types.Component {
	if t.Initializer != nil {
		t.Initializer(&t)
	}
	return &t
}

func (t *Text) Compose(ctx *types.Context) {
	ctx.Component = t
}

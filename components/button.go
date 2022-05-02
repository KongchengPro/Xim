package components

import "github.com/Re-Ch-Love/xim/types"

// Button 按钮
type Button struct {
	Content     string
	OnClick     func()
	Initializer func(*Button)
	types.BaseComponent
}

func (b Button) Render() (*types.RawComponent, bool) {
	return &types.RawComponent{
		LabelName:  "button",
		Attributes: nil,
		Content:    b.Content,
	}, false
}

func (b Button) Create() types.Component {
	if b.Initializer != nil {
		b.Initializer(&b)
	}
	return &b
}

func (b *Button) Compose(ctx *types.Context) {
	ctx.Component = b
	ctx.AddEventListener("click", b.OnClick)
}

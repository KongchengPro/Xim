package components

import (
	"github.com/Re-Ch-Love/xim/types"
)

//Panel 面板
type Panel struct {
	Color       string
	Children    []types.Component
	Initializer func(*Panel)
	types.BaseComponent
}

func (p Panel) Render() (*types.RawComponent, bool) {
	return &types.RawComponent{
		LabelName: "div",
		Attributes: map[string]string{
			"style": "background:" + p.Color + ";",
		},
	}, true
}

func (p Panel) Create() types.Component {
	if p.Initializer != nil {
		p.Initializer(&p)
	}
	return &p
}

func (p *Panel) Compose(ctx *types.Context) {
	ctx.Component = p
	ctx.AddChildren(p.Children)
}

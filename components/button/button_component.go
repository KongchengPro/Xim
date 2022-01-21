package button

import (
	"github.com/kongchengpro/xim/internal/utils"
	"github.com/kongchengpro/xim/types/callback"
	"github.com/kongchengpro/xim/types/component"
	"github.com/kongchengpro/xim/types/style"
)

// Button 按钮
type Button struct {
	id      string
	Name    string
	Content string
	OnClick callback.Func
}

func (c *Button) GetName() string {
	return c.Name
}

func (c *Button) GetCallbackMap() callback.Map {
	callbackMap := make(callback.Map)
	callbackMap["click"] = c.OnClick
	return callbackMap
}

func (c *Button) GetID() string {
	if c.id == "" {
		c.id = utils.GenerateID()
	}
	return c.id
}

func (c *Button) GetElementType() string {
	return "button"
}

func (c *Button) GetInnerHTML() string {
	return c.Content
}

func (c *Button) GetStyle() style.Style {
	return nil
}

func (c *Button) GetChildComponents() []component.Component {
	return nil
}

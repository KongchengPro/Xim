package button

import (
	"gitee.com/kogic/xim/internal/utils"
	"gitee.com/kogic/xim/types/callback"
	"gitee.com/kogic/xim/types/component"
	"gitee.com/kogic/xim/types/style"
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

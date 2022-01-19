package button

import (
	"gitee.com/kogic/xim/internal/utils"
	"gitee.com/kogic/xim/types"
)

// Button 按钮
type Button struct {
	id      string
	Name    string
	Content string
	OnClick func(args ...interface{})
}

func (c *Button) GetName() string {
	return c.Name
}

func (c *Button) GetCallbackMap() types.CallbackMap {
	callbackMap := make(types.CallbackMap)
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

func (c *Button) GetStyle() types.Style {
	return nil
}

func (c *Button) GetChildComponents() []types.Component {
	return nil
}

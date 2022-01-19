package text

import (
	"gitee.com/kogic/xim/internal/utils"
	"gitee.com/kogic/xim/types"
)

// Text 文本
type Text struct {
	id      string
	Name    string
	Style   *TextStyle
	Content string
}

func (c *Text) GetName() string {
	return c.Name
}

func (c *Text) GetCallbackMap() types.CallbackMap {
	return nil
}

func (c *Text) GetID() string {
	if c.id == "" {
		c.id = utils.GenerateID()
	}
	return c.id
}

func (c *Text) GetElementType() string {
	return "p"
}

func (c *Text) GetInnerHTML() string {
	return c.Content
}

func (c *Text) GetStyle() types.Style {
	return c.Style
}

func (c *Text) GetChildComponents() []types.Component {
	return nil
}

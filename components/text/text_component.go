package text

import (
	"github.com/kongchengpro/xim/internal/utils"
	"github.com/kongchengpro/xim/types/callback"
	"github.com/kongchengpro/xim/types/component"
	"github.com/kongchengpro/xim/types/style"
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

func (c *Text) GetCallbackMap() callback.Map {
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

func (c *Text) GetStyle() style.Style {
	return c.Style
}

func (c *Text) GetChildComponents() []component.Component {
	return nil
}

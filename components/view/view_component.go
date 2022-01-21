package view

import (
	"gitee.com/kogic/xim/internal/utils"
	"gitee.com/kogic/xim/types/callback"
	. "gitee.com/kogic/xim/types/component"
	. "gitee.com/kogic/xim/types/style"
)

// View 视图
// 仅作为承载Components的容器
type View struct {
	id         string
	Name       string
	Components []Component
}

func (c *View) GetName() string {
	return c.Name
}

func (c *View) GetCallbackMap() callback.Map {
	return nil
}

func (c *View) GetID() string {
	if c.id == "" {
		c.id = utils.GenerateID()
	}
	return c.id
}

func (c *View) GetElementType() string {
	return "div"
}

func (c *View) GetInnerHTML() string {
	return ""
}

func (c *View) GetStyle() Style {
	return nil
}

func (c *View) GetChildComponents() []Component {
	return c.Components
}

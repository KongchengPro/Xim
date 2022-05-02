package types

import "github.com/Re-Ch-Love/xim/internal/utils"

type Component interface {
	Id() string
	Compose(*Context)
	Create() Component
}

type BaseComponent struct {
	id string
}

func (c *BaseComponent) SetId(id string) {
	c.id = id
}

func (c *BaseComponent) Id() string {
	if c.id == "" {
		c.id = utils.GenerateID()
	}
	return c.id
}

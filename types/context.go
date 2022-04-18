package types

type Context struct {
	Component      Component
	EventListeners map[string]func()
	Children       []*Context
}

// AddEventListener adds an event listener to the context.
func (c *Context) AddEventListener(event string, callback func()) {
	if c.EventListeners == nil {
		c.EventListeners = make(map[string]func())
	}
	c.EventListeners[event] = callback
}

// Apply adds all the components in the array to the context
func (c *Context) Apply(comps []Component) {
	for _, childComp := range comps {
		childCtx := &Context{
			Component: childComp,
		}
		childComp.Compose(childCtx)
		c.Children = append(c.Children, childCtx)
	}
}

// Insert adds n child Component to the context
func (c *Context) Insert(childComp Component) {
	childCtx := &Context{
		Component: childComp,
	}
	childComp.Compose(childCtx)
	c.Children = append(c.Children, childCtx)
}

func (c *Context) Search(id string) Component {
	for _, child := range c.Children {
		if child.Component.Id() == id {
			return child.Component
		}
		if child.Children != nil {
			comp := child.Search(id)
			if comp != nil {
				return comp
			}
		}
	}
	return nil
}

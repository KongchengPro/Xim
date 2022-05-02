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

// AddChildren adds all the components in the array to the context
func (c *Context) AddChildren(comps []Component) {
	for _, childComp := range comps {
		childCtx := &Context{}
		childComp.Compose(childCtx)
		c.Children = append(c.Children, childCtx)
	}
}

// AddChild adds n child Component to the context
func (c *Context) AddChild(childComp Component) {
	childCtx := &Context{
		Component: childComp,
	}
	childComp.Compose(childCtx)
	c.Children = append(c.Children, childCtx)
}

// Find 用深度优先算法在Context.Children中搜索指定id的Component并返回它在每一层的索引列表和搜索到的Component对象
func (c *Context) Find(id string) ([]int, *Context) {
	var indexList []int
	var target *Context
	for i, child := range c.Children {
		indexList = append(indexList, i)
		if child.Component.Id() == id {
			target = child
			break
		}
		if child.Children != nil {
			indexList, target = child.Find(id)
			if target != nil {
				indexList = append([]int{i}, indexList...)
				break
			}
		}
	}
	return indexList, target
}

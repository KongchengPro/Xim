package types

//type Value = js.Value
//type Func func(this Value, args ...Value)
//type CallbackMap map[string]Func

type Value = any

type RawComponent struct {
	LabelName string
	Attribute map[string]string
	Content   string
	Children  []*RawComponent
}

type Event struct {
	Name string
	Data Value
}
type BaseComponent struct {
	eventChan chan *Event
}

func (c BaseComponent) Events() chan *Event {
	return c.eventChan
}

func (c BaseComponent) SetEvents(events chan *Event) {
	c.eventChan = events
}

type Context struct {
	Component Component
	Children  []*Context
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

// Insert adds a child Component to the context
func (c *Context) Insert(childComp Component) {
	childCtx := &Context{
		Component: childComp,
	}
	childComp.Compose(childCtx)
	c.Children = append(c.Children, childCtx)
}

type Component interface {
	Events() chan *Event
	SetEvents(chan *Event)
	Compose(*Context)
}

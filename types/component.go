package types

type Cs []Component

type CallbackMap map[string]func(args ...interface{})

type Component interface {
	GetID() string
	GetName() string
	GetElementType() string
	GetInnerHTML() string
	GetStyle() Style
	GetCallbackMap() CallbackMap
	GetChildComponents() []Component
}

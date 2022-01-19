package types

type Cs []Component

type Component interface {
	GetID() string
	GetElementType() string
	GetInnerHTML() string
	GetStyle() Style
	GetChildComponents() []Component
}

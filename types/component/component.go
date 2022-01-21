package component

import (
	"github.com/kongchengpro/xim/types/callback"
	"github.com/kongchengpro/xim/types/style"
)

type Cs []Component

type Component interface {
	GetID() string
	GetName() string
	GetElementType() string
	GetInnerHTML() string
	GetStyle() style.Style
	GetCallbackMap() callback.Map
	GetChildComponents() []Component
}

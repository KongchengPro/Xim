package component

import (
	"gitee.com/kogic/xim/types/callback"
	"gitee.com/kogic/xim/types/style"
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

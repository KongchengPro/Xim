package root

import (
	"gitee.com/kogic/xim/internal/constants"
	"gitee.com/kogic/xim/types/callback"
	"gitee.com/kogic/xim/types/component"
	"gitee.com/kogic/xim/types/style"
)

type Root struct {
}

func (c Root) GetName() string {
	return ""
}

func (r Root) GetCallbackMap() callback.Map {
	return nil
}

func (r Root) GetID() string {
	return constants.RootComponentID
}

func (r Root) GetElementType() string {
	return ""
}

func (r Root) GetInnerHTML() string {
	return ""
}

func (r Root) GetStyle() style.Style {
	return nil
}

func (r Root) GetChildComponents() []component.Component {
	return nil
}

package root

import (
	"gitee.com/kogic/xim/internal/constants"
	"gitee.com/kogic/xim/types"
)

type Root struct {
}

func (c Root) GetName() string {
	return ""
}

func (r Root) GetCallbackMap() types.CallbackMap {
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

func (r Root) GetStyle() types.Style {
	return nil
}

func (r Root) GetChildComponents() []types.Component {
	return nil
}

package root

import (
	"gitee.com/kogic/xim/internal/constants"
	"gitee.com/kogic/xim/types"
)

type Root struct {
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

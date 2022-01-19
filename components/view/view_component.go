package view

import (
	"gitee.com/kogic/xim/internal/utils"
	. "gitee.com/kogic/xim/types"
)

// View 视图
// 仅作为承载Components的容器
type View struct {
	id         string
	Components []Component
}

func (v *View) GetID() string {
	if v.id == "" {
		v.id = utils.GenerateID()
	}
	return v.id
}

func (v *View) GetElementType() string {
	return "div"
}

func (v *View) GetInnerHTML() string {
	return ""
}

func (v *View) GetStyle() Style {
	return nil
}

func (v *View) GetChildComponents() []Component {
	return v.Components
}

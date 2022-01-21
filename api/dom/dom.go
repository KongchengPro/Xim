package dom

import (
	"gitee.com/kogic/xim/types/component"
)

var ComponentMap = make(map[string]component.Component)

func GetComponentByPath(path string) component.Component {
	return ComponentMap[path]
}

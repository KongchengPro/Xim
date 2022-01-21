package dom

import (
	"github.com/kongchengpro/xim/types/component"
)

var ComponentMap = make(map[string]component.Component)

func GetComponentByPath(path string) component.Component {
	return ComponentMap[path]
}

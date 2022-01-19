package dom

import "gitee.com/kogic/xim/types"

var ComponentMap = make(map[string]types.Component)

func GetComponentByPath(path string) types.Component {
	return ComponentMap[path]
}

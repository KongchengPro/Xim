package counter

import (
	. "github.com/Re-Ch-Love/xim/components"
	"github.com/Re-Ch-Love/xim/types"
	"strconv"
)

type states struct {
	num types.MutableState[int]
}

var storage = &types.Storage[states]{
	&states{
		num: types.MutableStateOf[int](0),
	},
	map[string]func(*states, ...any){
		"increment": func(s *states, args ...any) {
			s.num.SetValue(s.num.Get() + 1)
		},
		"decrement": func(s *states, args ...any) {
			s.num.SetValue(s.num.Get() - 1)
		},
	},
}

var Counter = Panel{
	Color: "#fcfaed",
	Children: []types.Component{
		Text{
			Initializer: func(text *Text) {
				text.Content = types.NewDynamicData[string](func() string {
					return "Count: " + strconv.Itoa(storage.States.num.Value(text.Id()))
				})
			},
		}.Create(),
		Button{
			Content: "Increment",
			OnClick: func() {
				storage.Commit("increment")
			},
		}.Create(),
		Button{
			Content: "Decrement",
			OnClick: func() {
				storage.Commit("decrement")
			},
		}.Create(),
	},
}.Create()

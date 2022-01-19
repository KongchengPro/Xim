package main

import (
	"gitee.com/kogic/xim"
	. "gitee.com/kogic/xim/components/root"
	. "gitee.com/kogic/xim/components/text"
	. "gitee.com/kogic/xim/components/view"
	. "gitee.com/kogic/xim/types"
)

func main() {
	xim.SetTitle("Hello Xim")
	xim.Render(&View{
		Components: Cs{
			&Text{
				Content: "Hello Xim",
				Style: &TextStyle{
					FontSize: "40px",
				},
			},
		},
	}, Root{})
}

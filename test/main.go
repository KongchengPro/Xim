package main

import (
	"gitee.com/kogic/xim"
	. "gitee.com/kogic/xim/component"
	. "gitee.com/kogic/xim/style"
)

type Cs []Component

func main() {
	xim.Render(View{
		Components: Cs{
			&Text{
				Text: "Hello World",
				Style: Style{
					FontSize: "20px",
				},
			},
		},
	})
}

package main

import (
	"fmt"
	"gitee.com/kogic/xim"
	"gitee.com/kogic/xim/api/dom"
	. "gitee.com/kogic/xim/components/button"
	. "gitee.com/kogic/xim/components/root"
	. "gitee.com/kogic/xim/components/text"
	. "gitee.com/kogic/xim/components/view"
	"gitee.com/kogic/xim/types/callback"
	. "gitee.com/kogic/xim/types/component"
)

func main() {
	xim.SetTitle("Hello Xim")
	xim.Init(&View{
		Components: Cs{
			&Text{
				Name:    "MainText",
				Content: "Hello Xim",
				Style: &TextStyle{
					FontSize: "40px",
				},
			},
			&Button{
				Content: "Click Me",
				OnClick: func(this callback.Value, args ...callback.Value) {
					fmt.Printf("%#v\n%#v\n", this, args)
					c, ok := dom.GetComponentByPath("MainText").(*Text)
					if ok {
						c.Content = "Hello Kogic"
						xim.Refresh(c)
					}
				},
			},
		},
	}, Root{})
	select {}
}

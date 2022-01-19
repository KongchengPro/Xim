package main

import (
	"gitee.com/kogic/xim"
	"gitee.com/kogic/xim/api/dom"
	. "gitee.com/kogic/xim/components/button"
	. "gitee.com/kogic/xim/components/root"
	. "gitee.com/kogic/xim/components/text"
	. "gitee.com/kogic/xim/components/view"
	. "gitee.com/kogic/xim/types"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
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
				OnClick: func(args ...interface{}) {
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

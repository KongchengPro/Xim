package main

import (
	"github.com/LouisCheng-CN/xim"
	"github.com/LouisCheng-CN/xim/examples/HelloWorld/frontend/components"
)

func main() {
	xim.IsDebug = true
	xim.SetTitle("Hello Xim")
	xim.SetContent(&components.HelloWorld)
	select {}
}

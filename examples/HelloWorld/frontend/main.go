//go:build js || wasm

package main

import (
	"github.com/LouisCheng-CN/xim"
)

func main() {
	xim.SetTitle("Hello Xim")
	xim.SetContent(&HelloWorldComp)
	select {}
}

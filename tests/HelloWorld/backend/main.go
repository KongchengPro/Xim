package main

import "github.com/kongchengpro/xim/server"

// 需要在项目根目录下运行
func main() {
	server.StartStaticServer("./static")
}

package main

import "gitee.com/kogic/xim/server"

// 需要在项目根目录下运行
func main() {
	server.StartStaticServer("./static")
}

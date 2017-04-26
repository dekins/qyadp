package main

import (
	_ "dein.top/qyadp/routers"
	"github.com/astaxie/beego"
)

func main() {
	//设置静态文件路径
	beego.SetStaticPath("/static","dist/static")
	beego.Run()
}

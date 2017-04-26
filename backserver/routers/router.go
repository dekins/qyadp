package routers

import (
	"dein.top/qyadp/backserver/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}

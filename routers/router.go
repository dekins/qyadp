package routers

import (
	"dein.top/qyadp/qyadp/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}

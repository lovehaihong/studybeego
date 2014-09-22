package routers

import (
	"github.com/astaxie/beego"
	"studybeego/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}

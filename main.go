package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"studybeego/controllers"
	"studybeego/models"
)

func init() {
	// 注册数据库
	models.RegisterDB()
}

func main() {
	// 开户 ORM 调试模式
	orm.Debug = true

	// 自动建表
	orm.RunSyncdb("default", false, true) // 名称，是否自动重建，

	// 注册 beego 路由
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.Router("/reply", &controllers.ReplyController{})
	beego.Router("/reply/add", &controllers.ReplyController{}, "post:Add")
	beego.Router("/reply/delete", &controllers.ReplyController{}, "get:Delete")
	beego.AutoRouter(&controllers.TopicController{})

	//启动 beego
	beego.Run()
}

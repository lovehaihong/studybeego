package controllers

import (
	"github.com/astaxie/beego"
	"studybeego/models"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Data["IsHome"] = true
	this.TplNames = "home.html"

	this.Data["IsLogin"] = checkAccount(this.Ctx)

	topics, err := models.GetAllTopics(this.Input().Get("cate"), true)
	if err != nil {
		beego.Error(err)
	} else {
		this.Data["Topics"] = topics
	}

	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error()
	}

	this.Data["Categories"] = categories
}

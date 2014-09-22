package controllers

import (
	"github.com/astaxie/beego"
	"studybeego/models"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	opt := this.Input().Get("opt")

	switch opt {
	case "add":
		name := this.Input().Get("name")
		if len(name) == 0 {
			break
		}

		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}

		this.Redirect("/category", 301)
		return
	case "del":
		id := this.Input().Get("id")
		if len(id) == 0 {
			break
		}

		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}

		this.Redirect("/category", 301)
		return
	}

	this.Data["IsCategory"] = true
	this.TplNames = "category.html"

	var err error
	this.Data["Categories"], err = models.GetAllCategories()

	if err != nil {
		beego.Error(err)
	}
}

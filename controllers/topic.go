package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"studybeego/models"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsTopic"] = true
	this.TplNames = "topic.html"

	topics, err := models.GetAllTopics("", false)
	if err != nil {
		beego.Error(err)
	} else {
		this.Data["Topics"] = topics
	}
}

func (this *TopicController) Add() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.TplNames = "topic_add.html"

	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error()
	}

	this.Data["Categories"] = categories
}

func (this *TopicController) View() {
	this.TplNames = "topic_view.html"

	topic, err := models.GetTopic(this.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	this.Data["Topic"] = topic
	this.Data["Tid"] = this.Ctx.Input.Param("0")

	replies, err := models.GetAllReplies(this.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		return
	}

	this.Data["Replies"] = replies
	this.Data["IsLogin"] = checkAccount(this.Ctx)
}

func (this *TopicController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	title := this.Input().Get("title")
	content := this.Input().Get("content")
	category := this.Input().Get("category")
	tid := this.Input().Get("tid")

	fmt.Println(category)

	var err error
	if len(tid) == 0 {
		err = models.AddTopic(title, category, content)
		if err != nil {
			beego.Error(err)
		}
	} else {
		err = models.ModifyTopic(tid, title, category, content)
	}

	this.Redirect("/topic", 302)
}

func (this *TopicController) Modify() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	this.TplNames = "topic_modify.html"
	tid := this.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error()
	}

	this.Data["Categories"] = categories

	this.Data["Topic"] = topic
	this.Data["Tid"] = tid
}

func (this *TopicController) Delete() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	err := models.DeleteTopic(this.Input().Get("tid"))
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic", 302)
}

package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3" // 驱动注册，执行里面的init函数
	"os"
	"path"
	"strconv"
	"time"
)

const (
	_DB_NAME        = "data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Category struct {
	Id              int64     // 默认为主键
	Title           string    `orm:"null"`       // orm 默认为 255字节
	Created         time.Time `orm:"index;null"` // tag 为反射的时候使用
	Views           int64     `orm:"index;null"` // "index"意为建立索引
	TopicTime       time.Time `orm:"index;null"`
	TopicCount      int64     `orm:"null"`
	TopicLastUserId int64     `orm:"null"`
}

type Topic struct { // orm 模型
	Id              int64
	Uid             int64     `orm:"null"` // 谁写的
	Title           string    `orm:"null"` // 标题，默认为255字幕
	Category        string    `orm:"null"`
	Content         string    `orm:"size(5000);null"` // 内容，设为5000字节
	Attachment      string    `orm:"null"`            // 附件
	Created         time.Time `orm:"index;null"`      // tag 为反射的时候使用
	Updated         time.Time `orm:"index;null"`
	Views           int64     `orm:"index;null"`
	Author          string    `orm:"null"`
	ReplyTime       time.Time `orm:"null"`
	ReplyCount      int64     `orm:"null"`
	ReplyLastUserId int64     `orm:"null"`
}

type Comment struct {
	Id      int64
	Tid     int64     `orm:"null"`
	Name    string    `orm:"null"`
	Content string    `orm:"size(1000)"`
	Created time.Time `orm:"index;null"`
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) { // 检查数据库文件是否存在
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	orm.RegisterModel(new(Category), new(Topic), new(Comment))     // 注册模型，传入指针
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DR_Sqlite)             // 注册驱动，orm 允许操作N个数据库
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10) // 必须有一个数据库名称为"default"
}

func AddTopic(title, category, content string) error {
	o := orm.NewOrm()

	topic := &Topic{
		Title:     title,
		Category:  category,
		Content:   content,
		Created:   time.Now(),
		Updated:   time.Now(),
		ReplyTime: time.Now(),
	}

	_, err := o.Insert(topic)

	return err
}

func AddCategory(name string) error {
	o := orm.NewOrm()

	cate := &Category{Title: name}

	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}

	_, err = o.Insert(cate)
	if err != nil {
		return err
	}

	return nil
}

func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	cate := &Category{Id: cid}
	_, err = o.Delete(cate)

	return err
}

func GetAllTopics(cate string, isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)

	qs := o.QueryTable("topic")

	var err error
	if isDesc {
		if len(cate) > 0 {
			qs = qs.Filter("category", cate)
		}
		_, err = qs.OrderBy("-created").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}
	return topics, err
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)

	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}

func GetTopic(tid string) (*Topic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()
	topic := new(Topic)

	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, err
	}

	topic.Views++
	_, err = o.Update(topic)

	return topic, err
}

func ModifyTopic(tid, title, category, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}

	err = o.Read(topic)
	if err == nil { // 匹配原来已有的数据
		topic.Title = title
		topic.Category = category
		topic.Content = content
		topic.Updated = time.Now()
		o.Update(topic)
	}

	return err
}

func DeleteTopic(tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}
	o.Delete(topic)

	return err
}

func AddReply(tid, nickname, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	reply := &Comment{
		Tid:     tidNum,
		Name:    nickname,
		Content: content,
		Created: time.Now(),
	}

	o := orm.NewOrm()
	_, err = o.Insert(reply)
	return err
}

func GetAllReplies(tid string) (replies []*Comment, err error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	replies = make([]*Comment, 0)

	o := orm.NewOrm()
	qs := o.QueryTable("comment")
	_, err = qs.Filter("tid", tidNum).All(&replies)

	return replies, err
}

func DeleteReply(rid string) error {
	ridNum, err := strconv.ParseInt(rid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	reply := &Comment{Id: ridNum}
	o.Delete(reply)

	return err
}

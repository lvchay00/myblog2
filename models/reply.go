package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Reply struct {
	Id      int       `orm:"pk;auto"`
	Article *Article  `orm:"rel(fk)"`
	Email   string    `orm:"size(100)"`
	Content string    `orm:"type(text)"`
	User    *User     `orm:"rel(fk)"`
	InTime  time.Time `orm:"auto_now_add;type(datetime)"`
}

//通过id 查找回复
func FindReplyById(id int) Reply {
	o := orm.NewOrm()
	var reply Reply
	o.QueryTable(reply).RelatedSel("Article").Filter("Id", id).One(&reply)
	return reply
}

//通过文章查找回复
func FindReplyByArticle(article *Article) []*Reply {
	o := orm.NewOrm()
	var reply Reply
	var replies []*Reply
	o.QueryTable(reply).RelatedSel("Article").Filter("Article", article).OrderBy("-InTime").All(&replies)
	return replies
}

//通过用户查找回复
func FindReplyByUser(user *User, limit int) []*Reply {
	o := orm.NewOrm()
	var reply Reply
	var replies []*Reply
	o.QueryTable(reply).RelatedSel("Article", "User").Filter("User", user).OrderBy("-InTime").Limit(limit).All(&replies)
	return replies
}

//保存回复
func SaveReply(reply *Reply) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(reply)
	return id
}

//通过文章删除回复
func DeleteReplyByArticle(article *Article) {
	o := orm.NewOrm()
	var reply Reply
	var replies []Reply
	o.QueryTable(reply).Filter("Article", article).All(&replies)
	for _, reply := range replies {
		o.Delete(&reply)
	}
}

//删除回复
func DeleteReply(reply *Reply) {
	o := orm.NewOrm()
	o.Delete(reply)
}

//通过用户删除回复
func DeleteReplyByUser(user *User) {
	o := orm.NewOrm()
	o.Raw("delete form reply where user_id = ?", user.Id).Exec()
}

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

func FindReplyById(id int) Reply {
	o := orm.NewOrm()
	var reply Reply
	o.QueryTable(reply).RelatedSel("Article").Filter("Id", id).One(&reply)
	return reply
}

func FindReplyByArticle(article *Article) []*Reply {
	o := orm.NewOrm()
	var reply Reply
	var replies []*Reply
	o.QueryTable(reply).RelatedSel("Article").Filter("Article", article).OrderBy("-InTime").All(&replies)
	return replies
}

func FindReplyByUser(user *User, limit int) []*Reply {
	o := orm.NewOrm()
	var reply Reply
	var replies []*Reply
	o.QueryTable(reply).RelatedSel("Article", "User").Filter("User", user).OrderBy("-InTime").Limit(limit).All(&replies)
	return replies
}

func SaveReply(reply *Reply) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(reply)
	return id
}

func DeleteReplyByArticle(article *Article) {
	o := orm.NewOrm()
	var reply Reply
	var replies []Reply
	o.QueryTable(reply).Filter("Article", article).All(&replies)
	for _, reply := range replies {
		o.Delete(&reply)
	}
}

func DeleteReply(reply *Reply) {
	o := orm.NewOrm()
	o.Delete(reply)
}

func DeleteReplyByUser(user *User) {
	o := orm.NewOrm()
	o.Raw("delete form reply where user_id = ?", user.Id).Exec()
}

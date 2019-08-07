package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/lvchay00/myblog2/filters"
	"github.com/lvchay00/myblog2/models"
)

type ReplyController struct {
	beego.Controller
}

func (c *ReplyController) Save() {
	content := c.Input().Get("content")
	email := c.Input().Get("email")
	if len(content) == 0 {
		c.Ctx.WriteString("回复内容不能为空")
	} else {
		tid, _ := strconv.Atoi(c.Input().Get("tid"))
		if tid == 0 {
			c.Ctx.WriteString("回复的文章不存在")
		} else {
			_, user := filters.IsLogin(c.Ctx)
			if user.Id == 0 {
				user.Id = 1
				user.Username = "游客"
			}
			article := models.FindArticleById(tid)
			reply := models.Reply{Content: content, Article: &article, User: &user, Email: email}
			models.SaveReply(&reply)
			models.IncrReplyCount(&article)
			c.Redirect("/article/"+strconv.Itoa(tid), 302)
		}
	}
}

func (c *ReplyController) Delete() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if id > 0 {
		reply := models.FindReplyById(id)
		tid := reply.Article.Id
		models.ReduceReplyCount(reply.Article)
		models.DeleteReply(&reply)
		c.Redirect("/article/"+strconv.Itoa(tid), 302)
	} else {
		c.Ctx.WriteString("回复不存在")
	}
}

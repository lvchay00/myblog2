package controllers

import (
	"fmt"
	"strconv"

	"github.com/lvchay00/myblog2/filters"
	"github.com/lvchay00/myblog2/models"
	"github.com/lvchay00/myblog2/utils"

	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
	times int8
}

//首页
func (c *IndexController) Index() {
	c.Data["PageTitle"] = "首页"
	//找到所有分区
	c.Data["Sections"] = models.FindAllSection()
	//是否为登录
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Controller.Ctx)
	//从路径中获取当前页面
	p, _ := strconv.Atoi(c.Ctx.Input.Query("p"))
	if p == 0 {
		p = 1
	}
	//从路径中获取是否跳转页面
	prv, _ := strconv.Atoi(c.Ctx.Input.Query("prv"))
	fmt.Println("prv", prv)
	//从配置中获取 页面的大小
	size, _ := beego.AppConfig.Int("page.size")

	//从路径中获取当前分区
	s, _ := strconv.Atoi(c.Ctx.Input.Query("s"))
	c.Data["S"] = s

	total, articles := models.FindArticleByArticleSectionsUsingSection(p, size, s)
	//生成概览，只显前400个字
	for i, article := range articles {

		if len([]rune(article.Content)) > 400 {
			articles[i].Content = string([]rune(article.Content)[:350])
		} else {
			articles[i].Content = string([]rune(article.Content))
		}
	}
	c.Data["articles"] = articles

	//生成路径
	url := fmt.Sprintf("/?s=%d&p=", s)

	//生成分页标签
	c.Data["pagebar"] = utils.PaginationToString(p, int(total), size, 5, 10, url)

	c.Layout = "layout/layout.html"
	c.TplName = "index.html"
}

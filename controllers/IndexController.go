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

	c.Data["Sections"] = models.FindAllSection()

	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Controller.Ctx)

	p, _ := strconv.Atoi(c.Ctx.Input.Query("p"))
	if p == 0 {
		p = 1
	}

	prv, _ := strconv.Atoi(c.Ctx.Input.Query("prv"))
	fmt.Println("prv", prv)

	size, _ := beego.AppConfig.Int("page.size")

	s, _ := strconv.Atoi(c.Ctx.Input.Query("s"))
	c.Data["S"] = s

	total, articles := models.FindArticleByArticleSectionsUsingSection(p, size, s)

	for i, article := range articles {

		if len([]rune(article.Content)) > 400 {
			articles[i].Content = string([]rune(article.Content)[:350])
		} else {
			articles[i].Content = string([]rune(article.Content))
		}
	}
	c.Data["articles"] = articles

	url := fmt.Sprintf("/?s=%d&p=", s)

	c.Data["pagebar"] = utils.PaginationToString(p, int(total), size, 5, 10, url)

	c.Layout = "layout/layout.html"
	c.TplName = "index.html"
}

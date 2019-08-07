package controllers

import (
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/lvchay00/myblog2/filters"
	"github.com/lvchay00/myblog2/models"
)

type ArticleController struct {
	beego.Controller
}

//上传文件接口
func (c *ArticleController) Uploade() {
	req := c.Ctx.Request
	req.ParseForm()
	// 接收文件
	uploadFile, handle, err := req.FormFile("file")
	if err != nil {
		c.Data["output"] = "接收文件失败"
		return
	}
	filepath := c.Input().Get("path")
	filepath1 := "./" + filepath + "/"

	// 保存文件
	os.Mkdir(filepath1, 0777)
	saveFile, err := os.OpenFile(filepath1+handle.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		c.Data["output"] = "文件创建失败"
		return
	}
	_, err = io.Copy(saveFile, uploadFile)
	if err != nil {
		c.Data["output"] = "文件创建失败"
		return
	}

	defer uploadFile.Close()
	defer saveFile.Close()
	// 检查图片后缀
	ext := strings.ToLower(path.Ext(handle.Filename))
	if ext == ".jpg" || ext == ".png" {
		c.Data["output"] = filepath + "/" + handle.Filename + "上传成功  "
		// 上传图片成功

	} else {
		c.Data["output"] = filepath + "/" + handle.Filename + "上传成功  "
	}
	c.Layout = "layout/layout.html"
	c.TplName = "ArticleController/create.html"
	return
}
func (c *ArticleController) Create() {
	beego.ReadFromRequest(&c.Controller)
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Controller.Ctx)
	c.Data["PageTitle"] = "发布文章"
	section := models.FindAllSection()
	section = append(section, &models.Section{Id: 0, Name: ""})
	c.Data["Sections"] = section
	c.Layout = "layout/layout.html"
	c.TplName = "ArticleController/create.html"
}

//删除重复项
func DeleteDuplicateNumbers(a []int, l int) (s []int) {
	var i, j, t int
	if l <= 0 {
		return
	}
	for j = 0; j < l; j++ {
		t = a[j]
		for i = j + 1; i < l; i++ {
			if t == a[i] { /*判断是否有相同的*/
				a[i] = 0
			}
		}
	}
	return a
}
func (c *ArticleController) Save() {
	flash := beego.NewFlash()
	title, content, sid := c.Input().Get("title"), c.Input().Get("content"), c.Input().Get("sid")
	sid1, sid2, sid3 := c.Input().Get("sid1"), c.Input().Get("sid2"), c.Input().Get("sid3")

	fmt.Println("sid", sid, title)
	if len(title) == 0 || len(title) > 120 {
		flash.Error("文章标题不能为空且不能超过120个字符")
		flash.Store(&c.Controller)
		c.Redirect("/article/create", 302)
	} else if len(sid) == 0 {
		flash.Error("请选择文章版块")
		flash.Store(&c.Controller)
		c.Redirect("/article/create", 302)
	} else {
		var a []int
		s, err := strconv.ParseInt(sid, 0, 8)
		if err == nil {
			a = append(a, int(s))
		}
		s1, err := strconv.ParseInt(sid1, 0, 8)
		if err == nil {
			a = append(a, int(s1))
		}
		s2, err := strconv.ParseInt(sid2, 0, 8)
		if err == nil {
			a = append(a, int(s2))
		}
		s3, err := strconv.ParseInt(sid3, 0, 8)
		if err == nil {
			a = append(a, int(s3))
		}
		a = DeleteDuplicateNumbers(a, 4)

		_, user := filters.IsLogin(c.Ctx)
		article := models.Article{Title: title, Content: content, User: &user}
		err, id := models.SaveArticle(&article)

		if err != nil {
			flash.Error("标题重复或者格式错误")
			flash.Store(&c.Controller)
			c.Redirect("/article/create", 302)
		} else {
			for i := 0; i < len(a); i++ {
				if a[i] != 0 {
					models.SaveArticleSections(id, int64(a[i]))
				}
			}
			c.Redirect("/article/"+strconv.FormatInt(id, 10), 302)
		}
	}
}

func (c *ArticleController) Detail() {
	id := c.Ctx.Input.Param(":id")
	tid, _ := strconv.Atoi(id)
	if tid > 0 {
		c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Controller.Ctx)
		article := models.FindArticleById(tid)
		models.IncrView(&article) //查看+1
		c.Data["PageTitle"] = article.Title
		c.Data["Article"] = article
		c.Data["Replies"] = models.FindReplyByArticle(&article)
		c.Layout = "layout/layout.html"
		c.TplName = "ArticleController/detail.html"
	} else {
		c.Ctx.WriteString("文章不存在")
	}
}

func (c *ArticleController) Edit() {
	beego.ReadFromRequest(&c.Controller)
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if id > 0 {
		article := models.FindArticleById(id)
		c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Controller.Ctx)
		c.Data["PageTitle"] = "编辑文章"
		//用于下拉控件选项
		c.Data["Sections"] = models.FindAllSection()
		c.Data["Article"] = article
		article.Section = models.FindSectionByArticle(int64(id))

		if len(article.Section) > 0 {
			c.Data["Section0"] = article.Section[0].Id
		} else {
			c.Data["Section0"] = 0
		}
		if len(article.Section) > 1 {
			c.Data["Section1"] = article.Section[1].Id
		} else {
			c.Data["Section1"] = 0
		}
		if len(article.Section) > 2 {
			c.Data["Section2"] = article.Section[2].Id
		} else {
			c.Data["Section2"] = 0
		}
		if len(article.Section) > 3 {
			c.Data["Section3"] = article.Section[3].Id
		} else {
			c.Data["Section3"] = 0
		}
		c.Layout = "layout/layout.html"
		c.TplName = "ArticleController/edit.html"
	} else {
		c.Ctx.WriteString("文章不存在")
	}
}

func (c *ArticleController) Update() {
	flash := beego.NewFlash()
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	title, content, sid := c.Input().Get("title"), c.Input().Get("content"), c.Input().Get("sid")
	sid1, sid2, sid3 := c.Input().Get("sid1"), c.Input().Get("sid2"), c.Input().Get("sid3")

	if len(title) == 0 || len(title) > 120 {
		flash.Error("文章标题不能为空且不能超过120个字符")
		flash.Store(&c.Controller)
		c.Redirect("/article/edit/"+strconv.Itoa(id), 302)
	} else if len(sid) == 0 && len(sid1) == 0 && len(sid2) == 0 && len(sid3) == 0 {
		flash.Error("请选择文章分类")
		flash.Store(&c.Controller)
		c.Redirect("/article/edit/"+strconv.Itoa(id), 302)
	} else {
		var a []int
		s, err := strconv.ParseInt(sid, 0, 8)
		if err == nil {
			a = append(a, int(s))
		}
		s1, err := strconv.ParseInt(sid1, 0, 8)
		if err == nil {
			a = append(a, int(s1))
		}
		s2, err := strconv.ParseInt(sid2, 0, 8)
		if err == nil {
			a = append(a, int(s2))
		}
		s3, err := strconv.ParseInt(sid3, 0, 8)
		if err == nil {
			a = append(a, int(s3))
		}
		a = DeleteDuplicateNumbers(a, len(a))
		models.DeleteArticleSectionByArticle(int64(id))
		for i := 0; i < len(a); i++ {
			if a[i] != 0 {
				models.SaveArticleSections(int64(id), int64(a[i]))
			}
		}
		article := models.FindArticleById(id)
		article.Title = title
		article.Content = content
		models.UpdateArticle(&article)
		c.Redirect("/article/"+strconv.Itoa(id), 302)
	}
}

func (c *ArticleController) Delete() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if id > 0 {
		article := models.FindArticleById(id)
		models.DeleteArticle(&article)

		c.Redirect("/", 302)
	} else {
		c.Ctx.WriteString("文章不存在")
	}
}

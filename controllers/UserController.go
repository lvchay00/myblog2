package controllers

import (
	"github.com/astaxie/beego"
	"github.com/lvchay00/myblog2/filters"
	"github.com/lvchay00/myblog2/models"
	"github.com/lvchay00/myblog2/utils"
	"github.com/sluu99/uuid"
)

type UserController struct {
	beego.Controller
}

//登录页 /login  get
func (c *UserController) LoginPage() {
	IsLogin, _ := filters.IsLogin(c.Ctx)
	if IsLogin {
		c.Redirect("/", 302)
	} else {
		beego.ReadFromRequest(&c.Controller)

		c.Data["PageTitle"] = "登录"
		c.Layout = "layout/layout.html"
		c.TplName = "login.tpl"
	}
}

//验证登录 /login post
func (c *UserController) Login() {
	flash := beego.NewFlash()
	username, password := c.Input().Get("username"), c.Input().Get("password")

	if flag, user := models.Login(username, utils.Md5([]byte(password))); flag {
		filters.SetUserCookie(c.Controller.Ctx, user)
		c.Redirect("/", 302)
	} else {
		flash.Error("用户名或密码错误")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
	}
}

//注册页 /register  get
func (c *UserController) RegisterPage() {
	IsLogin, _ := filters.IsLogin(c.Ctx)
	if IsLogin {
		c.Redirect("/", 302)
	} else {
		beego.ReadFromRequest(&c.Controller)
		c.Data["PageTitle"] = "注册"
		c.Layout = "layout/layout.html"
		c.TplName = "register.tpl"
	}
}

//验证注册 /register post
func (c *UserController) Register() {
	flash := beego.NewFlash()
	username, password := c.Input().Get("username"), c.Input().Get("password")
	if len(username) == 0 || len(password) == 0 {
		flash.Error("用户名或密码不能为空")
		flash.Store(&c.Controller)
		c.Redirect("/register", 302)
	} else if flag, _ := models.FindUserByUserName(username); flag {
		flash.Error("用户名已被注册")
		flash.Store(&c.Controller)
		c.Redirect("/register", 302)
	} else {
		//使用UUID 生成token
		var token = uuid.Rand().Hex()
		user := models.User{Username: username, Password: utils.Md5([]byte(password)), Token: token}
		models.SaveUser(&user)
		filters.SetUserCookie(c.Controller.Ctx, user)
		c.Redirect("/", 302)
	}
}

//退出  /loginout
func (c *UserController) Logout() {
	filters.DeleteCookie(c.Controller.Ctx)
	c.Redirect("/", 302)
}

//关于 /about
func (c *UserController) About() {
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Controller.Ctx)
	c.Data["PageTitle"] = "关于"
	c.Layout = "layout/layout.html"
	c.TplName = "about.tpl"
}

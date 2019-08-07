package routers

import (
	"github.com/astaxie/beego"
	"github.com/lvchay00/myblog2/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{}, "GET:Index")
	beego.Router("/login", &controllers.UserController{}, "GET:LoginPage")
	beego.Router("/login", &controllers.UserController{}, "POST:Login")
	beego.Router("/register", &controllers.UserController{}, "GET:RegisterPage")
	beego.Router("/register", &controllers.UserController{}, "POST:Register")
	beego.Router("/logout", &controllers.UserController{}, "GET:Logout")
	beego.Router("/about", &controllers.UserController{}, "GET:About")

	beego.Router("/article/create", &controllers.ArticleController{}, "GET:Create")
	beego.Router("/article/create", &controllers.ArticleController{}, "POST:Save")
	beego.Router("/article/create/Uploade", &controllers.ArticleController{}, "Post:Uploade")

	beego.Router("/article/:id([0-9]+)", &controllers.ArticleController{}, "GET:Detail")
	beego.Router("/article/edit/:id([0-9]+)", &controllers.ArticleController{}, "GET:Edit")
	beego.Router("/article/edit/:id([0-9]+)", &controllers.ArticleController{}, "POST:Update")

	beego.Router("/article/delete/:id([0-9]+)", &controllers.ArticleController{}, "GET:Delete")

	beego.Router("/reply/save", &controllers.ReplyController{}, "POST:Save")
	beego.Router("/reply/delete/:id([0-9]+)", &controllers.ReplyController{}, "GET:Delete")

	beego.Router("/AddClass", &controllers.AddClassController{}, "GET:Index")
	beego.Router("/AddClass/Add", &controllers.AddClassController{}, "Post:Add")
	beego.Router("/AddClass/Delete", &controllers.AddClassController{}, "Post:Delete")

	beego.Router("/up", &controllers.FileController{}, "GET:UploadHandle")         // 上传
	beego.Router("/up", &controllers.FileController{}, "Post:UploadHandle")        // 上传
	beego.Router("/uploaded/", &controllers.FileController{}, "GET:ShowPicHandle") //显示图片
}

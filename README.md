#                  Blog简单介绍
## 1.演示说明
 项目演示地址：http://104.224.163.70:8081/
 
 测试账号：
 
 用户名：test
 
 密码：test
 
 登录账号后 允许创建文章，编辑文章，删除文章

## 2.项目简单介绍

  使用 beego 的MVC 结构
  
  用户注册和登录使用cookie 和uuid.
  
  数据库使用mysql

### 控制层    路由

```
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
```

### 模型层
   数据库结构 ER图
 ![image](https://github.com/lvchay00/myblog2/blob/master/static/imgs/%E6%95%B0%E6%8D%AE%E5%BA%93ER%E5%9B%BE.png?raw=true)  
 
###  视图层 
   使用 bootstartup3
   
   富文本编辑器 wangEditor-2

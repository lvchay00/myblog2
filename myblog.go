package main

import (
	"fmt"

	"os"

	"github.com/lvchay00/myblog2/fileserver"
	"github.com/lvchay00/myblog2/models"
	_ "github.com/lvchay00/myblog2/routers"
	_ "github.com/lvchay00/myblog2/templates"
	_ "github.com/lvchay00/myblog2/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	url := beego.AppConfig.String("jdbc.url")
	fmt.Println(url)
	port := beego.AppConfig.String("jdbc.port")
	username := beego.AppConfig.String("jdbc.username")
	password := beego.AppConfig.String("jdbc.password")

	orm.RegisterModel(
		new(models.User),
		new(models.Article),
		new(models.Section),
		new(models.Reply))

	orm.RegisterDataBase("default", "mysql", username+":"+password+"@tcp("+url+":"+port+")/myblog?charset=utf8", 30)

	orm.RunSyncdb("default", false, true)
}

func main() {

	fmt.Println("PID:", os.Getpid())
	go fileserver.File_server()

	beego.SetLogger("file", `{"filename":"LOG.log"}`)
	beego.Run()
}

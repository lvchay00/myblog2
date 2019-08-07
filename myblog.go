package main

import (
	"fmt"

	"os"
	_ "runtime"

	//	"time"

	"github.com/lvchay00/myblog2/fileserver"
	"github.com/lvchay00/myblog2/models"
	_ "github.com/lvchay00/myblog2/routers"
	_ "github.com/lvchay00/myblog2/templates"
	_ "github.com/lvchay00/myblog2/utils"

	//	"syscall"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	//	"github.com/zserge/webview"
)

// var (
// 	user32           = syscall.NewLazyDLL("User32.dll")
// 	getSystemMetrics = user32.NewProc("GetSystemMetrics")
// )

// func GetSystemMetrics(nIndex int) int {
// 	index := uintptr(nIndex)
// 	ret, _, _ := getSystemMetrics.Call(index)
// 	return int(ret)
// }

// const (
// 	SM_CXSCREEN = 0
// 	SM_CYSCREEN = 1
// )

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

// func display() {
// 	x := GetSystemMetrics(SM_CXSCREEN)
// 	y := GetSystemMetrics(SM_CYSCREEN)
// 	time.Sleep(10 * time.Microsecond)
// 	// Open wikipedia in a 800x600 resizable window
// 	// webview.Open("Minimal webview example",
// 	// 	"http://localhost:8080", 800, 600, true)
// 	webview.Open("Minimal webview example",
// 		"http://localhost:8081", x, y-40, true)
// }

func main() {
	//	go display()
	//runtime.GOMAXPROCS(1)

	fmt.Println("PID:", os.Getpid())
	go fileserver.File_server()

	// if beego.AppConfig.String("runmode") == "dev" {
	// 	orm.Debug = true
	// }
	// orm.Debug = true

	beego.SetLogger("file", `{"filename":"LOG.log"}`)
	beego.Run()
}

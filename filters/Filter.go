package filters

import (
	"strings"

	"github.com/astaxie/beego/context"
	"github.com/lvchay00/myblog2/models"
)

func IsLogin(ctx *context.Context) (bool, models.User) {
	arr := strings.Split(ctx.GetCookie("auth"), "|")
	var user models.User
	if len(arr) == 2 {
		token := arr[0]
		flag, user := models.FindUserByToken(token)
		return flag, user
	}
	return false, user
}
func SetUserCookie(ctx *context.Context, user models.User) {
	s := strings.Split(ctx.Request.RemoteAddr, ":")
	ctx.SetCookie("auth", user.Token+"|"+s[0], 7*86400)
}
func DeleteCookie(ctx *context.Context) {
	ctx.SetCookie("auth", "")
}

// var FilterUser = func(ctx *context.Context) {
// 	ok, _ := IsLogin(ctx)
// 	if !ok {
// 		ctx.Redirect(302, "/login")
// 	}
// }

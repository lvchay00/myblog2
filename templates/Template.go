package utils

import (
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/russross/blackfriday"
	"github.com/xeonx/timeago"
)

func FormatTime(time time.Time) string {
	return timeago.Chinese.Format(time)
}

func NoHtml(str string) string {
	return strings.Replace(strings.Replace(str, "<script", "&lt;script", -1), "script>", "script&gt;", -1)
}

func Markdown(content string) string {
	return string(blackfriday.MarkdownCommon([]byte(NoHtml(content))))
}

func init() {
	beego.AddFuncMap("timeago", FormatTime)
	beego.AddFuncMap("markdown", Markdown)
}

package utils

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"math"
	"net/url"
	"strings"
)

// type Pager struct {
// 	Page int //当前页数

// 	Totalnum int //总条数

// 	Pagesize int //页大小
// 	offset   int //偏移页数
// 	linknum  int //显示的总页数
// 	url      string
// }

func PaginationToString(page, totalnum, pagesize, offset, linknum int, url string) string {
	if totalnum <= pagesize {
		return ""
	}

	var buf bytes.Buffer
	var from, to, totalpage int

	// offset   //偏移页数为5
	//linknum //显示的页数

	totalpage = int(math.Ceil(float64(totalnum) / float64(pagesize)))

	//查找开始和结束位置
	if totalpage < linknum {
		from = 1       //开始位置
		to = totalpage //结束位置
	} else {
		from = page - offset //当前页数 - 偏移页数
		to = from + linknum
		if from < 1 {
			from = 1
			to = from + linknum - 1
		} else if to > totalpage {
			to = totalpage
			from = totalpage - linknum + 1
		}
	}

	//前一页
	buf.WriteString("<div class=\"container\"><ul class=\"pagination\">\n")

	if page > 1 {
		buf.WriteString(fmt.Sprintf("<li class=\"page-item\"><a class=\"page-link\" href=\"%s%d&prv=1\">Previous</a></li>\n", url, (page - 1)))
	} else {
		buf.WriteString(fmt.Sprintf("<li class=\"page-item\"><a class=\"page-link\" href=\"%s%d&prv=1\">Previous</a></li>\n", url, 1))
	}

	//中间
	for i := from; i <= to; i++ {
		if i == page {
			buf.WriteString(fmt.Sprintf("<li class=\"page-item active\"><a class=\"page-link\" href=\"%s%d\">%d</a></li>\n", url, i, i))
		} else {
			buf.WriteString(fmt.Sprintf("<li class=\"page-item\"><a class=\"page-link\" href=\"%s%d\">%d</a></li>\n", url, i, i))
		}
	}

	//后一页
	if page < totalpage {
		buf.WriteString(fmt.Sprintf("<li class=\"page-item\"><a class=\"page-link\" href=\"%s%d&prv=2\">Next</a></li>\n", url, (page + 1)))
	} else {
		buf.WriteString(fmt.Sprintf("<li class=\"page-item\"><a class=\"page-link\" href=\"%s%d&prv=2\">Next</a></li>\n", url, totalpage))
	}

	buf.WriteString("</ul></div>")
	fmt.Println(buf.String())
	return buf.String()
}
func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func Rawurlencode(str string) string {
	return strings.Replace(url.QueryEscape(str), "+", "%20", -1)
}

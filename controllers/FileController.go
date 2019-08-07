package controllers

import (
	_ "errors"
	"fmt"
	"io"
	"io/ioutil"
	_ "net/http"
	"os"
	"path"
	"strings"

	"github.com/astaxie/beego"
)

type FileController struct {
	beego.Controller
}

//上传图像接口
func (c *FileController) UploadHandle() {
	w := c.Ctx.ResponseWriter
	req := c.Ctx.Request
	req.ParseForm()
	if req.Method != "POST" {
		w.Write([]byte(html))
	} else {
		// 接收文件
		uploadFile, handle, err := req.FormFile("file")
		if err != nil {
			w.Write([]byte(err.Error()))
			fmt.Fprintf(w, "接收文件失败")
			return
		}
		filepath := c.Input().Get("path")
		filepath1 := "./" + filepath + "/"

		// 保存文件 可能会创建相同的文件夹 ，会报错误，不影响使用
		err = os.Mkdir(filepath1, 0777)
		// if err != nil {
		// 	w.Write([]byte(err.Error()))
		// 	fmt.Fprintf(w, filepath + "文件创建失败")
		// 	return
		// }
		saveFile, err := os.OpenFile(filepath1+handle.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			w.Write([]byte(err.Error()))
			fmt.Fprintf(w, filepath+"文件创建失败")
			return
		}
		_, err = io.Copy(saveFile, uploadFile)
		if err != nil {
			w.Write([]byte(err.Error()))
			fmt.Fprintf(w, filepath+"文件拷贝失败")
			return
		}
		defer uploadFile.Close()
		defer saveFile.Close()
		// 检查图片后缀
		ext := strings.ToLower(path.Ext(handle.Filename))
		if ext == ".jpg" || ext == ".png" {
			//	fmt.Fprintf(w, filepath + "/"+handle.Filename +"上传成功  ")
			// 上传图片成功
			w.Write([]byte("<a  target='_blank' href='/" + filepath + "/" + handle.Filename + "'>" + handle.Filename + "</a>"))
		} else {
			fmt.Fprintf(w, filepath+"/"+handle.Filename+"上传成功  ")
		}

	}
	return
}

// 显示图片接口
func (c *FileController) ShowPicHandle() {
	w := c.Ctx.ResponseWriter
	req := c.Ctx.Request
	file, err := os.Open("." + req.URL.Path)
	if err != nil {
		w.Write([]byte(err.Error()))
		fmt.Fprintf(w, "文件打开失败")
		return
	}
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	buff, err := ioutil.ReadAll(file)
	if err != nil {
		w.Write([]byte(err.Error()))
		fmt.Fprintf(w, "读取文件失败")
		return
	}
	w.Write(buff)
	if err != nil {
		fmt.Println(err)
	}
}

const html = `<html>
    <head></head>
    <body>
        <form method="post" enctype="multipart/form-data">
		    文件目录：<input type="text" name="path" value="static/blog" />
			<p>  </p>
            <input type="file" name="file" />
            <input type="submit" />
        </form>
    </body>
</html>`

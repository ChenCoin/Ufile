package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func uploadFile(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		io.WriteString(w, "<html><head><title>上传</title></head>"+
			"<body><form action='#' method=\"post\" enctype=\"multipart/form-data\">"+
			"<label>上传图片</label>"+":"+
			"<input type=\"file\" name='file'  /><br/><br/>    "+
			"<label><input type=\"submit\" value=\"上传图片\"/></label></form></body></html>")
	} else {
		//获取文件内容 要这样获取
		file, _, err := r.FormFile("file")
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "")
			return
		}
		defer file.Close()
		//创建文件

		r.ParseForm()
		dscPath := strings.Join(r.Form["path"], "")
		fW, err := os.Create("." + dscPath)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "文件创建失败")
			return
		}
		defer fW.Close()
		_, err = io.Copy(fW, file)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "文件保存失败")
			return
		}
		fmt.Fprintf(w, "success")
	}
}
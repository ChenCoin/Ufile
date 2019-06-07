package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		_, _ = io.WriteString(w, "<html><head><title>上传</title></head>"+
			"<body><form action='#' method=\"post\" enctype=\"multipart/form-data\">"+
			"<label>上传文件</label>"+":"+
			"<input type=\"file\" name='file'  /><br/><br/>    "+
			"<label><input type=\"submit\" value=\"上传文件\"/></label></form></body></html>")
	} else {
		file, _, err := r.FormFile("file")
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "")
			return
		}
		defer file.Close()

		err2 := r.ParseForm()
		if err2 != nil {
			w.WriteHeader(404)
			fmt.Fprintf(w, "")
			return
		}

		dscPath := strings.Join(r.Form["path"], "")
		fW, err := os.Create("." + dscPath)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "文件创建失败")
			return
		}
		defer fW.Close()
		
		fmt.Fprintf(w, "save file : " + file)
		_, err = io.Copy(fW, file)
		fmt.Fprintf(w, "save file suc : " + file)
		
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "文件保存失败")
			return
		}
		fmt.Fprintf(w, "success")
	}
}

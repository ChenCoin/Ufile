package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func checkPath(path string) bool {
	realPath, err := filepath.Abs(path)
	if err != nil{
		return false
	}
	return path == realPath
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(strings.Join(r.Form["method"], ""))
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Wrold!")
}

type FileType struct {
	Name string `json:"name"`
	Size int64 `json:"size"`
	Mode os.FileMode `json:"mode"`
	ModTime time.Time `json:"time"`
	IsDir bool `json:"isDir"`
}

func getDir(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	path := strings.Join(r.Form["path"], "")
	fmt.Println("path:", path)
	if !checkPath(path) {
		w.WriteHeader(403)
		fmt.Fprintf(w, "")
		return
	}
	path = "." + path

	files, err1 := ioutil.ReadDir(path)
	if err1 != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, "")
		return
	}

	var fileData []FileType
	for _, f := range files {
		fileData = append(fileData, FileType{f.Name(), f.Size(), f.Mode(), f.ModTime(), f.IsDir()})
	}

	jsonStr, err2 := json.Marshal(fileData)
	if err2 != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "")
	}else{
		fmt.Fprintf(w, string(jsonStr))
	}
}

func getFile(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	filePath := strings.Join(r.Form["path"], "")
	fmt.Println("download:", filePath)
	if !checkPath(filePath) {
		w.WriteHeader(403)
		fmt.Fprintf(w, "")
		return
	}
	filePath = "." + filePath

	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, "")
	}

	fmt.Println("download form ", file.Name())
	w.Header().Add("Content-Disposition", "attachment; filename=" + path.Base(file.Name()))
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("Content-Transfer-Encoding", "binary")
	http.ServeContent(w, r, file.Name(), time.Now(), file)
	fmt.Fprintf(w, "file")
}

func copyFile(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	srcPath := strings.Join(r.Form["old"], "")
	dstPath := strings.Join(r.Form["new"], "")
	fmt.Println("copy form ", srcPath, " to ", dstPath)
	if !checkPath(srcPath) || !checkPath(dstPath) {
		w.WriteHeader(403)
		fmt.Fprintf(w, "")
		return
	}
	srcPath = "." + srcPath
	dstPath = "." + dstPath

	f, err := os.Stat(srcPath)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "")
		return
	}
	if f.IsDir() {
		fmt.Fprintf(w, "can't copy dir")
		return
	}

	src, err := os.Open(srcPath)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "")
		return
	}
	defer src.Close()

	dst, err := os.Create(dstPath)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "")
		return
	}
	defer dst.Close()

	_, err2 := io.Copy(dst, src)
	if err2 != nil{
		w.WriteHeader(500)
		fmt.Fprintf(w, "")
		return
	}

	fmt.Fprintf(w, "success")
}

func cutFile(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	srcPath := strings.Join(r.Form["old"], "")
	dstPath := strings.Join(r.Form["new"], "")
	fmt.Println("cut form ", srcPath, " to ", dstPath)
	if !checkPath(srcPath) || !checkPath(dstPath) {
		w.WriteHeader(403)
		fmt.Fprintf(w, "")
		return
	}
	srcPath = "." + srcPath
	dstPath = "." + dstPath

	err := os.Rename(srcPath,dstPath)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "")
	}else{
		fmt.Fprintf(w, "success")
	}
}

func deleteFile(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	path := strings.Join(r.Form["path"], "")
	fmt.Println("delete ", path)
	if !checkPath(path) {
		w.WriteHeader(403)
		fmt.Fprintf(w, "")
		return
	}
	path = "." + path

	err := os.RemoveAll(path)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "")
	}else{
		fmt.Fprintf(w, "success")
	}
}

func renameFile(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	srcPath := strings.Join(r.Form["old"], "")
	dstPath := strings.Join(r.Form["new"], "")
	fmt.Println("rename form ", srcPath, " to ", dstPath)
	if !checkPath(srcPath) || !checkPath(dstPath) {
		w.WriteHeader(403)
		fmt.Fprintf(w, "")
		return
	}
	srcPath = "." + srcPath
	dstPath = "." + dstPath

	err := os.Rename(srcPath,dstPath)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "")
	}else{
		fmt.Fprintf(w, "success")
	}
}

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

func createDir(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	path := strings.Join(r.Form["path"], "")
	fmt.Println("create dir ", path)
	if !checkPath(path) {
		w.WriteHeader(403)
		fmt.Fprintf(w, "")
		return
	}
	path = "." + path

	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "")
	}else{
		fmt.Fprintf(w, "success")
	}
}

func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "index")
}

func main() {
	Check()
	http.HandleFunc("/get/dir/", getDir)
	http.HandleFunc("/get/file/", getFile)
	http.HandleFunc("/put/path/", copyFile)
	http.HandleFunc("/post/file/", cutFile)
	http.HandleFunc("/delete/file/", deleteFile)
	http.HandleFunc("/post/name/", renameFile)
	http.HandleFunc("/put/file/", uploadFile)
	http.HandleFunc("/put/dir/", createDir)
	http.HandleFunc("/hello/", sayhelloName)
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

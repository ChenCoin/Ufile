package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func copyFiles(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, "")
		return
	}

	srcPath := strings.Join(r.Form["old"], "")
	dstPath := strings.Join(r.Form["new"], "")
	fmt.Println("copy form ", srcPath, " to ", dstPath)
	if !check(srcPath) || !check(dstPath) {
		w.WriteHeader(403)
		_, _ = fmt.Fprintf(w, "")
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

	result := false
	if f.IsDir() {
		result = copyDir(srcPath, dstPath)
	} else {
		result = copyFile(srcPath, dstPath)
	}
	if result {
		fmt.Fprintf(w, "success")
	} else {
		w.WriteHeader(500)
		fmt.Fprintf(w, "")
	}
}

func copyFile(srcPath, dstPath string) bool {
	src, err := os.Open(srcPath)
	if err != nil {
		fmt.Println("err1 :", err)
		return false
	}
	defer src.Close()
	dst, err := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("err1 :", err)
		return false
	}
	defer dst.Close()
	_, err2 := io.Copy(dst, src)
	if err2 != nil {
		fmt.Println("err1 :", err2)
		return false
	}
	return true
}

func copyDir(src, dst string) bool {
	if os.Mkdir(dst, os.ModePerm) != nil {
		return false
	}
	result := true
	files, err := ioutil.ReadDir(src)
	if err != nil {
		return false
	}
	for _, f := range files {
		if f.Name() == filepath.Base(src) {
			continue
		}
		if f.IsDir() {
			if !copyDir(src+"/"+f.Name(), dst+"/"+f.Name()) {
				result = false
				break
			}
		} else if !copyFile(src+"/"+f.Name(), dst+"/"+f.Name()) {
			result = false
			break
		}
		fmt.Println(f.Name())
	}
	return result
}

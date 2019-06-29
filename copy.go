package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// it can't be named copy due to duplicate name
func copyFiles(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(404)
		_, _ = w.Write([]byte("404"))
		log.Printf("rename ParseForm 404")
		return
	}

	srcPath := r.URL.Path
	dstPath := strings.Join(r.Form["new"], "")
	if !check(srcPath) || !check(dstPath) {
		w.WriteHeader(404)
		_, _ = w.Write([]byte("404"))
		log.Printf("copy %s %s 404", srcPath, dstPath)
		return
	}
	srcPath = "." + srcPath
	dstPath = "." + dstPath

	f, err := os.Stat(srcPath)
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte("500"))
		log.Printf("copy %s %s 500", srcPath, dstPath)
		return
	}

	result := false
	if f.IsDir() {
		result = copyDir(srcPath, dstPath)
	} else {
		result = copyFile(srcPath, dstPath)
	}
	if !result {
		w.WriteHeader(500)
		_, _ = w.Write([]byte("500"))
		log.Printf("copy %s %s 500", srcPath, dstPath)
	} else {
		_, _ = w.Write([]byte("success"))
		log.Printf("copy %s %s success", srcPath, dstPath)
	}
}

func copyFile(srcPath, dstPath string) bool {
	src, err := os.Open(srcPath)
	if err != nil {
		log.Printf("err : %s", err)
		return false
	}
	defer src.Close()
	dst, err := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Printf("err : %s", err)
		return false
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	if err != nil {
		log.Printf("err : %s", err)
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
		log.Printf("err : %s", err)
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
	}
	return result
}

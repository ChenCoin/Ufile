package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func copyFile(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
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
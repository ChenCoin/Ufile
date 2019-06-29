package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func rename(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(404)
		_, _ = w.Write([]byte("404"))
		log.Printf("rename ParseForm 404")
		return
	}

	srcPath := r.URL.Path
	dstPath := strings.Join(r.Form["to"], "")

	if !check(srcPath) || !check(dstPath) {
		w.WriteHeader(404)
		_, _ = w.Write([]byte("404"))
		log.Printf("rename %s %s 404", srcPath, dstPath)
		return
	}

	err = os.Rename("."+srcPath, "."+dstPath)
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte("500"))
		log.Printf("rename %s %s 500", srcPath, dstPath)
	} else {
		_, _ = w.Write([]byte("success"))
		log.Printf("rename %s %s success", srcPath, dstPath)
	}
}

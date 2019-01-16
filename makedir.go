package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func createDir(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, "")
	}

	path := strings.Join(r.Form["path"], "")
	fmt.Println("create dir ", path)
	if !check(path) {
		w.WriteHeader(403)
		fmt.Fprintf(w, "")
		return
	}
	path = "." + path

	err2 := os.Mkdir(path, os.ModePerm)
	if err2 != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "")
	} else {
		fmt.Fprintf(w, "success")
	}
}

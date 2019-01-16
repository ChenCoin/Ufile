package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func renameFile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, "")
	}
	srcPath := strings.Join(r.Form["old"], "")
	dstPath := strings.Join(r.Form["new"], "")
	fmt.Println("rename form ", srcPath, " to ", dstPath)
	if !check(srcPath) || !check(dstPath) {
		w.WriteHeader(403)
		fmt.Fprintf(w, "")
		return
	}
	srcPath = "." + srcPath
	dstPath = "." + dstPath

	err2 := os.Rename(srcPath, dstPath)
	if err2 != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "")
	} else {
		fmt.Fprintf(w, "success")
	}
}

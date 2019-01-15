package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

func getFile(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, "")
		return
	}
	filePath := strings.Join(r.Form["path"], "")
	fmt.Println("download:", filePath)
	if !check(filePath) {
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
		return
	}

	fmt.Println("download form ", file.Name())
	w.Header().Add("Content-Disposition", "attachment; filename=" + path.Base(file.Name()))
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("Content-Transfer-Encoding", "binary")
	http.ServeContent(w, r, file.Name(), time.Now(), file)
	fmt.Fprintf(w, "file")
}
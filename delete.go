package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func deleteFile(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err!=nil{
		w.WriteHeader(404)
		fmt.Fprintf(w, "")
	}
	path := strings.Join(r.Form["path"], "")
	fmt.Println("delete ", path)
	if !check(path) {
		w.WriteHeader(403)
		fmt.Fprintf(w, "")
		return
	}
	path = "." + path

	err = os.RemoveAll(path)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "")
	}else{ fmt.Fprintf(w, "success") }
}
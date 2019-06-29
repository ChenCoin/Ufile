package main

import (
	"log"
	"net/http"
	"os"
)

func mkdir(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if !check(path) {
		w.WriteHeader(404)
		_, _ = w.Write([]byte("404"))
		log.Printf("mkdir %s 404", path)
		return
	}
	err := os.Mkdir("."+path, os.ModePerm)
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte("500"))
		log.Printf("mkdir %s 500", path)
	} else {
		_, _ = w.Write([]byte("success"))
		log.Printf("mkdir %s success", path)
	}
}

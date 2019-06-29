package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type FileType struct {
	Name    string      `json:"name"`
	Size    int64       `json:"size"`
	Mode    os.FileMode `json:"mode"`
	ModTime time.Time   `json:"time"`
	IsDir   bool        `json:"isDir"`
}

func list(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if !check(path) {
		w.WriteHeader(404)
		_, _ = w.Write([]byte("404"))
		log.Printf("list %s 404", path)
		return
	}

	path = "." + path

	files, err := ioutil.ReadDir(path)
	if err != nil {
		w.WriteHeader(404)
		_, _ = w.Write([]byte("404"))
		log.Printf("list.ReadDir %s 404", path)
		return
	}

	var fileData []FileType
	for _, f := range files {
		if f.Name() == filepath.Base(path) {
			continue
		}
		fileData = append(fileData, FileType{f.Name(), f.Size(), f.Mode(), f.ModTime(), f.IsDir()})
	}

	jsonStr, err := json.Marshal(fileData)
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte("500"))
		log.Printf("list %s 500", path)
	} else {
		_, _ = w.Write([]byte(string(jsonStr)))
		log.Printf("list %s success", path)
	}
}

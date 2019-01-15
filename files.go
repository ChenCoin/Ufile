package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type FileType struct {
	Name string `json:"name"`
	Size int64 `json:"size"`
	Mode os.FileMode `json:"mode"`
	ModTime time.Time `json:"time"`
	IsDir bool `json:"isDir"`
}

func getDir(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, "")
		return
	}
	path := strings.Join(r.Form["path"], "")
	fmt.Println("path:", path)
	if !check(path) {
		w.WriteHeader(403)
		fmt.Fprintf(w, "")
		return
	}
	path = "." + path

	files, err1 := ioutil.ReadDir(path)
	if err1 != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, "")
		return
	}

	var fileData []FileType
	for _, f := range files {
		fileData = append(fileData, FileType{f.Name(), f.Size(), f.Mode(), f.ModTime(), f.IsDir()})
	}

	jsonStr, err2 := json.Marshal(fileData)
	if err2 != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "")
	} else { fmt.Fprintf(w, string(jsonStr)) }
}
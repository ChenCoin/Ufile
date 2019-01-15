package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "index")
}

func main() {
	http.HandleFunc("/get/dir/", getDir)
	http.HandleFunc("/get/file/", getFile)
	http.HandleFunc("/put/path/", copyFile)
	http.HandleFunc("/post/file/", cutFile)
	http.HandleFunc("/delete/file/", deleteFile)
	http.HandleFunc("/post/name/", renameFile)
	http.HandleFunc("/put/file/", uploadFile)
	http.HandleFunc("/put/dir/", createDir)
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Println("error when create server")
	}
}
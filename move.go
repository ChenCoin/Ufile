package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func cutFile(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil{
		w.WriteHeader(404)
		fmt.Fprintf(w, "")
		return
	}
	srcPath := strings.Join(r.Form["old"], "")
	dstPath := strings.Join(r.Form["new"], "")
	fmt.Println("cut form ", srcPath, " to ", dstPath)
	if !check(srcPath) || !check(dstPath) {
		w.WriteHeader(403)
		fmt.Fprintf(w, "")
		return
	}
	srcPath = "." + srcPath
	dstPath = "." + dstPath

	err = os.Rename(srcPath,dstPath)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "")
	}else{
		fmt.Fprintf(w, "success")
	}
}
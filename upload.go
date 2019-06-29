package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		_, _ = io.WriteString(w, `
<html>
<head>
    <title>upload</title>
</head>
<body>
	<form enctype="multipart/form-data" action="#" method="post">
		<input type="file" name="files" />
		<input type="submit" value="upload" />
	</form>
</body>
</html>
`)
	} else {
		path := r.URL.Path
		if !check(path) {
			w.WriteHeader(404)
			_, _ = w.Write([]byte("404"))
			log.Printf("delete %s 404", path)
			return
		}

		err := r.ParseMultipartForm(1024000)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		m := r.MultipartForm
		files := m.File["files"]
		for i, _ := range files {
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			dst, err := os.Create("." + path + "/" + files[i].Filename)
			defer dst.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if _, err := io.Copy(dst, file); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			_, _ = w.Write([]byte("success"))
		}
	}
}

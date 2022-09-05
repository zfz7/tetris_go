package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
)

//go:embed build
var frontendFiles embed.FS

func main() {
	fmt.Println("Starting Server")
	http.Handle("/", http.FileServer(getFileSystem()))
	http.ListenAndServe(":8080", nil)
}

func getFileSystem() http.FileSystem {
	fsys, err := fs.Sub(frontendFiles, "build")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}
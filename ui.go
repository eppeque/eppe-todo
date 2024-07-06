package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

//go:embed ui/build/*
var folder embed.FS

func assets() (fs.FS, error) {
	return fs.Sub(folder, "ui/build")
}

func assignUIHandler() {
	ui, err := assets()

	if err != nil {
		log.Fatalln(err)
	}

	handler := http.FileServer(http.FS(ui))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			handler.ServeHTTP(w, r)
			return
		}

		f, err := ui.Open(strings.TrimPrefix(path.Clean(r.URL.Path), "/"))

		if err == nil {
			defer f.Close()
		}

		if os.IsNotExist(err) {
			r.URL.Path += ".html"
		}

		handler.ServeHTTP(w, r)
	})
}

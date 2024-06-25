package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed ui/*
var folder embed.FS

func assets() (fs.FS, error) {
	return fs.Sub(folder, "ui")
}

func assignUIHandler() {
	ui, err := assets()

	if err != nil {
		log.Fatalln(err)
	}

	handler := http.FileServer(http.FS(ui))
	http.Handle("/", http.StripPrefix("/", handler))
}

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/eppeque/todo-server/infra"
)

func handleInterrupt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-c
		infra.CloseDatabase()
		os.Exit(0)
	}()
}

func readFlag() int {
	port := flag.Int("p", 8080, "The port number to listen to")
	flag.Parse()
	return *port
}

func initData() {
	if err := infra.InitData(); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	initData()

	port := readFlag()
	addr := fmt.Sprintf(":%d", port)

	handleInterrupt()
	assignAPIHandlers()
	assignUIHandler()

	log.Printf("Listening on port %d...\n", port)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalln(err)
	}
}

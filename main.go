package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/eppeque/todo-server/models"
)

func handleInterrupt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("Shutting down...")
		os.Exit(0)
	}()
}

func main() {
	config := models.InitConfig()
	addr := fmt.Sprintf(":%d", config.Port)

	handleInterrupt()
	assignAPIHandlers(config)
	assignUIHandler()

	log.Printf("Listening on port %d...\n", config.Port)
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		log.Fatalln(err)
	}
}

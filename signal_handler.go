package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.Signal(0xa))
	for syscall.Signal(0xa) == <-sigs {
		log.Print("Recieved 0xa, reloading config")
		LoadConfig()
	}
}

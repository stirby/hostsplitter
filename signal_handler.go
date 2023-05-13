package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func SignalHandler() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.Signal(0xa))
	for {
		if <-sigs == syscall.Signal(0xa) {
			log.Print("Recieved 0xa, reloading config")
			LoadConfig()
		}
	}
}

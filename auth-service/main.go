package main

import (
	"log"

	"github.com/dilshodforever/nasiya-savdo/server"
)

func main() {
	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}
}

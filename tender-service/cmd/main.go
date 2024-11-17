package main

import (
	"github.com/dilshodforever/tender/internal/app"
	"github.com/dilshodforever/tender/internal/pkg/config"
)

func main() {
	config := config.Load()

	app.Run(&config)
}

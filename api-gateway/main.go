package main

import (
	"github.com/dilshodforever/tender/internal/app"
	"github.com/dilshodforever/tender/internal/pkg/config"
)

func main() {
	cf := config.Load()
	app.Run(cf)
}

package main

import (
	"github.com/mariosoaresreis/go-banking-hexagonal/app"
	"github.com/mariosoaresreis/go-banking-hexagonal/logger"
)

func main() {
	logger.Info("Iniciando banking")
	app.Start()
}

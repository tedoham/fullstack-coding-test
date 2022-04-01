package main

import (
	"github.com/tedoham/fullstack-coding-test/app"
	"github.com/tedoham/fullstack-coding-test/logger"
)

func main() {
	logger.Info("Starting the application")
	app.Start()
}

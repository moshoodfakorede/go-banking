package main

import (
	"github.com/Fakorede/banking/app"
	"github.com/Fakorede/banking/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}

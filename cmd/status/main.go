package main

import (
	"os"
	"time"

	"github.com/utkarsh-josh/hdfc/inithandler"
)

func main() {
	logger := inithandler.InitLogging()
	service := inithandler.InitService(logger, 20)
	go service.StatusChecker()
	err := inithandler.HandleRequests(logger, service)
	if err != nil {
		logger.Log(
			"method", "main",
			"msg", "Failed to listen to the endpoints",
			"err", err.Error(),
			"time", time.Now(),
		)
		os.Exit(1)
	}
}

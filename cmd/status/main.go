package main

import (
	"os"
	"time"

	"github.com/utkarsh-josh/hdfc/inithandler"
	"github.com/utkarsh-josh/hdfc/svcconst"
)

func main() {
	logger := inithandler.InitLogging(svcconst.ServiceLoggerPort)
	service := inithandler.InitService(logger, svcconst.PeriodicStatusCheckInterval)
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

package testutil

import (
	"github.com/go-kit/kit/log"

	"github.com/utkarsh-josh/hdfc/inithandler"
	"github.com/utkarsh-josh/hdfc/sdk"
	"github.com/utkarsh-josh/hdfc/svcconst"
)

// TestUtil will encapsulate all the components,sdk etc. required to run CT
type TestUtil struct {
	Logger log.Logger
	API    sdk.Service
}

// InitTestInfra initializes common components to bring service up
func InitTestInfra() *TestUtil {
	logger := inithandler.InitLogging(svcconst.TestEnvLoggerPort)
	service := inithandler.InitService(logger, svcconst.PeriodicStatusCheckInterval)
	go service.StatusChecker()
	go inithandler.HandleRequests(logger, service)
	api := sdk.NewServiceSDK(logger, svcconst.ServiceHostName)

	return &TestUtil{
		Logger: logger,
		API:    api,
	}
}

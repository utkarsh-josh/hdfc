package inithandler

import (
	"github.com/go-kit/kit/log"

	"github.com/utkarsh-josh/hdfc/service"
	serviceDL "github.com/utkarsh-josh/hdfc/service/dl"
	"github.com/utkarsh-josh/hdfc/svcconst"
	wp "github.com/utkarsh-josh/hdfc/workerpool"
)

// InitService initializes service bl layer and dl layer
func InitService(logger log.Logger, duration int) *service.BL {
	dl := serviceDL.NewDL(logger)
	workerpool := wp.NewWorkerPool(svcconst.MaxWorkerCount)
	bl := service.NewService(logger, workerpool, dl, duration)
	return bl
}

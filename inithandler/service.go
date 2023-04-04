package inithandler

import (
	"github.com/go-kit/kit/log"

	serviceDL "github.com/utkarsh-josh/hdfc/service/dl"
	"github.com/utkarsh-josh/hdfc/service"
)

func InitService(logger log.Logger, duration int) *service.BL {
	dl := serviceDL.NewDL(logger)
	bl := service.NewService(logger, dl, duration)
	return bl
}

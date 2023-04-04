package inithandler

import (
	"github.com/go-kit/kit/log"

	"github.com/utkarsh-josh/hdfc/service"
	serviceDL "github.com/utkarsh-josh/hdfc/service/dl"
)

// InitService initializes service bl layer and dl layer
func InitService(logger log.Logger, duration int) *service.BL {
	dl := serviceDL.NewDL(logger)
	bl := service.NewService(logger, dl, duration)
	return bl
}

package inithandler

import (
	"os"

	"github.com/go-kit/kit/log"
)

func InitLogging() log.Logger {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", "8080", "caller", log.DefaultCaller)

	return logger
}

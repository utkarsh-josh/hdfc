package inithandler

import (
	"os"

	"github.com/go-kit/kit/log"
)

// InitLogging initializes logging handler
func InitLogging(port string) log.Logger {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", port, "caller", log.DefaultCaller)

	return logger
}

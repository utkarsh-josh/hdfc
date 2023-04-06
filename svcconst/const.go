package svcconst

// LoggerPorts...
const (
	ServiceLoggerPort = "8080"
	TestEnvLoggerPort = "8081"
)

// Intervals...
const (
	PeriodicStatusCheckInterval   = 60
	TestEnvWaitTimeForStatusCheck = 70
)

// ServiceHostName...
const (
	ServiceHostName = "localhost:12000"
)

// Status...
const (
	StatusNotYetChecked = "StatusNotYetChecked"
	StatusUp            = "UP"
	StatusDown          = "DOWN"
)

// MaxWorkerCount...
const (
	MaxWorkerCount = 5
)

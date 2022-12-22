package constants

import "time"

const (
	GracefulTimeoutDuration = 30

	API = "api"

	WORKER = "worker"

	REQUEST_TIMEOUT = 2000 * time.Millisecond
)

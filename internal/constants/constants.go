package constants

import "time"

const (
	GracefulTimeoutDuration = 30

	API = "api"

	REQUEST_TIMEOUT = 2000 * time.Millisecond

	API_Key_Threshold = 97

	MigrationDir = "dir"

	DefaultMigrationDir = "migrations"

	Development = "dev"

	Env = "env"

	DefaultBasePath = "."
)

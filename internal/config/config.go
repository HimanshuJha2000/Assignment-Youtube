package config

import "fmt"

const (
	// FilePath - relative path to the config directory
	FilePath = "%s/config/%s"

	// EnvFilename - Filename format of env specific config file
	EnvFilename = "env.%s.toml"
)

var (
	// config : this will hold all the application configuration
	config appConfig
)

// appConfig global configuration struct definition
type appConfig struct {
	Database    Database    `toml:"database"`
	Application application `toml:"application"`
	Worker      worker      `toml:"worker"`
}

func LoadConfig(basePath string, env string) {
	fmt.Println("Loading environment for " + env)

	// reading env file and override config values; default env is dev
	LoadConfigFromFile(FilePath, basePath, EnvFilename, &config, env)
}

// GetConfig : will give the struct as value so that the actual config doesn't get tampered
func GetConfig() appConfig {
	return config
}

package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"strconv"
)

// loadConfigFromFile: load/overwrite config values from given basepath, filename and env
func LoadConfigFromFile(
	filePath string,
	basePath string,
	filename string,
	configStruct interface{},
	env string) {
	path := getFilePath(filePath, basePath, filename, env)

	_, statErr := os.Stat(path)

	if statErr != nil {
		panic("config file not found")
	}

	if _, err := toml.DecodeFile(path, configStruct); err != nil {
		panic("Invalid data type in configuration")
	}
}

// getFilePath: gives the file path based on the environment provided
// file path will be relative to the application and determined by basePath
func getFilePath(filePath string, basePath string, fileName string, env string) string {
	envFile := env

	if env != "" {
		fileName = fmt.Sprintf(fileName, envFile)
	}

	path := fmt.Sprintf(filePath, basePath, fileName)

	return path
}

func GetYoutubeURLRequestEndpoint(base_url string, api_key string, count int, query string) string {
	return base_url + "key=" + api_key + "&type=video&part=snippet&maxResults=" + strconv.Itoa(count) + "&order=date&q=" + query
}

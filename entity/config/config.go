package config

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

var config *Config

type Config struct {
	ProtectedBranches                          []string
	AutomaticPushWhenCommiting                 bool
	LengthBranchPrefix                         int
	AutomaticCheckoutToNewBranch               bool
	DefaultBranch                              string
	IgnoreValidations                          bool
	ChechoutDefaultBranchBeforeCreateNewBranch bool
}

func init() {
	CreateConfigFileIfNotExist()
}

func GetDefaultConfig() *Config {
	return &Config{
		ProtectedBranches:                          []string{"main", "release", "master"},
		AutomaticPushWhenCommiting:                 false,
		LengthBranchPrefix:                         6,
		AutomaticCheckoutToNewBranch:               true,
		DefaultBranch:                              "release",
		IgnoreValidations:                          false,
		ChechoutDefaultBranchBeforeCreateNewBranch: true,
	}

}

const (
	FILE_NAME        = "giteconfig.json"
	PATH_ENVIRONMENT = "Path"
	DEFAULT_PATH_WIN = "C:\\bin;"
)

func CreateConfigFileIfNotExist() {

	//setPathEnvironment()

	if _, err := os.Stat(FILE_NAME); errors.Is(err, os.ErrNotExist) {

		json, _ := json.MarshalIndent(GetDefaultConfig(), "", "\t")

		err = os.WriteFile(FILE_NAME, []byte(json), 777)
		if err != nil {
			panic("Não foi possivel criar o arquivo de configuração.")
		}

	}
}

func GetConfig() *Config {
	file, err := os.Open(FILE_NAME)

	if err != nil {
		return GetDefaultConfig()
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	json.Unmarshal(byteValue, &config)

	return config
}

// func setPathEnvironment() {
// 	path := os.Getenv(PATH_ENVIRONMENT)

// 	if !strings.Contains(path, DEFAULT_PATH_WIN) {
// 		var LAST_ELEMENT = len(path) - 1

// 		if path[LAST_ELEMENT] == ';' {
// 			path += DEFAULT_PATH_WIN
// 		} else {
// 			path += ";" + DEFAULT_PATH_WIN
// 		}

// 		if err := os.Setenv(PATH_ENVIRONMENT, path); err != nil {
// 			fmt.Println("Não foi possivel setar a variavel de ambiente PATH.\n Adicione manualmente o caminho para o exe ao fim da Variavel Path.")
// 		}
// 	}

// }

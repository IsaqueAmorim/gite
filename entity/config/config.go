package config

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

type Config struct {
	ProtectedBranches []string
	AutomaticPush     bool
	LengthPrefix      int
}

func init() {
	CreateConfigFileIfNotExist()
}

func GetDefaultConfig() *Config {
	return &Config{
		ProtectedBranches: []string{"main", "release", "master"},
		AutomaticPush:     false,
		LengthPrefix:      6,
	}
}

func CreateConfigFileIfNotExist() {

	if _, err := os.Stat("giteconfig.json"); errors.Is(err, os.ErrNotExist) {

		json, _ := json.Marshal(GetDefaultConfig())

		err = os.WriteFile("giteconfig.json", []byte(json), 777)
		if err != nil {
			panic("Não foi possivel criar o arquivo de configuração.")
		}

	}
}

func GetConfig() *Config {
	file, err := os.Open("giteconfig.json")

	if err != nil {
		return GetDefaultConfig()
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	var result *Config

	json.Unmarshal(byteValue, &result)

	return result
}

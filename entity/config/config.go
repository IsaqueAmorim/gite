package config

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"os/exec"
)

var config *Config

type Config struct {
	ProtectedBranches []string
	AutomaticPush     bool
	LengthPrefix      int
}

func init() {
	W()
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

	os.Setenv("GITE_PATH", "C:/bin/")

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

	json.Unmarshal(byteValue, &config)

	return config
}

func W() {

	exec.Command("Set-Item", "-Path", `Env:\PATH`, "-Value", `$env:PATH;"C:\bin"`)

}

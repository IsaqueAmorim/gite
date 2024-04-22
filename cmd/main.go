package main

import "github.com/IsaqueAmorim/gite/entity/config"

const (
	ENV_CONFIG_PATH      = "GITE_CONFIG_PATH"
	WINDOWS_DEFAULT_PATH = "C:/gite/"
)

func main() {

	//path := os.Getenv(ENV_CONFIG_PATH)
	config.GetConfig()

}

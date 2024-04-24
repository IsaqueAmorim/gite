package main

import (
	"github.com/IsaqueAmorim/gite/cli"
	"github.com/IsaqueAmorim/gite/entity/config"
	"github.com/spf13/cobra"
)

const (
	ENV_CONFIG_PATH      = "GITE_CONFIG_PATH"
	WINDOWS_DEFAULT_PATH = "C:/gite/"
)

func main() {

	var rootCmd = &cobra.Command{}
	config.GetConfig()

	rootCmd.AddCommand(cli.CommitCmd(), cli.BranchCmd())
	rootCmd.Execute()

}

package main

import (
	"fmt"

	"github.com/IsaqueAmorim/gite/entity/branch"
	"github.com/IsaqueAmorim/gite/entity/config"
	"github.com/spf13/cobra"
)

const (
	ENV_CONFIG_PATH      = "GITE_CONFIG_PATH"
	WINDOWS_DEFAULT_PATH = "C:/gite/"
)

func main() {

	config.GetConfig()
	var prefix, branchName string

	var cmd = &cobra.Command{
		Use:   "gite",
		Short: "Gite is a tool to help you with your git flow.",
		Run: func(cmd *cobra.Command, args []string) {
			b := branch.NewBranch(prefix, branchName)

			fmt.Print(b.GetFullBranchName())

		},
	}

	cmd.Flags().StringVarP(&prefix, "prefix", "p", "", "Prefix to be used in the branch name.")
	cmd.Flags().StringVarP(&branchName, "branch", "b", "", "Branch name.")

	cmd.Execute()

}

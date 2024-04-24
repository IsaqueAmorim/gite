package cli

import (
	"fmt"

	"github.com/IsaqueAmorim/gite/entity/branch"
	"github.com/spf13/cobra"
)

func BranchCmd() *cobra.Command {
	var (
		prefix, branchName string
	)

	commitCmd := &cobra.Command{
		Use:   "branch",
		Short: "Gite is a tool to help you with your git flow.",
		Run: func(cmd *cobra.Command, args []string) {
			b := branch.NewBranch(prefix, branchName)

			fmt.Print(b.GetFullBranchName())

		},
	}
	commitCmd.Flags().StringVarP(&branchName, "branch", "b", "", "Nome da Branch.")
	commitCmd.Flags().StringVarP(&prefix, "prefix", "p", "", "Prefixo para nome da Branch.")

	return commitCmd
}

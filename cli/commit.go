package cli

import (
	"github.com/IsaqueAmorim/gite/command"
	"github.com/IsaqueAmorim/gite/entity/branch"
	"github.com/IsaqueAmorim/gite/entity/commit"
	"github.com/IsaqueAmorim/gite/entity/config"
	"github.com/spf13/cobra"
)

func CommitCmd() *cobra.Command {
	var (
		Commit commit.Commit
		Config config.Config
	)

	commitCmd := &cobra.Command{
		Use:   "commit",
		Short: "Usado para fazer commit das alterações locais no branch atual.",
		Run: func(cmd *cobra.Command, args []string) {
			command.Commit(&Config, &Commit)
		},
	}
	commitCmd.Flags().StringVarP(&Commit.Message, "message", "m", "", "Mensagem do commit.")
	commitCmd.Flags().StringVarP(&Commit.Status, "status-refs", "r", "refs", "Status refs.")
	commitCmd.Flags().StringVarP(&Commit.Status, "status-coloses", "c", "closes", "Status closes.")
	commitCmd.Flags().StringVar(&Commit.Time, "time", "0h0min", "Tempo Gasto.")
	commitCmd.Flags().StringSliceVarP(&Commit.Ticket, "tickets", "t", []string{branch.GetCurrentBranch()}, "Tickets.")
	commitCmd.Flags().BoolVar(&Config.IgnoreValidations, "ignore-validations", false, "Ignorar Validações.")
	commitCmd.Flags().BoolVarP(&Config.AutomaticPushWhenCommiting, "automatic-push", "p", false, "Push Automático.")

	return commitCmd
}

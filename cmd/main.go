package main

import "github.com/spf13/cobra"

func main() {
}

func Commit() *cobra.Command {
	return &cobra.Command{
		Use:   "commit",
		Short: "Record changes to the repository",
		Long:  "Record changes to the repository",
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}
}

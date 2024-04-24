package command

import (
	"fmt"
	"slices"

	"github.com/IsaqueAmorim/gite/entity/branch"
	"github.com/IsaqueAmorim/gite/entity/commit"
	"github.com/IsaqueAmorim/gite/entity/config"
)

func Commit(config *config.Config, commit *commit.Commit) {

	branch := branch.GetCurrentBranch()

	if !slices.Contains(config.ProtectedBranches, branch) {

		GitAddAll()
		GitCommitWithMessage(commit.Message, "0h0min")
		gitPush(config.AutomaticPushWhenCommiting, branch)
	} else {

		fmt.Println("Você não pode fazer commit em um branch protegido!")
	}

}

func GitAddAll() {
	//exec.Command("git", "add", "--all")
	fmt.Println("git add --all")
}

func GitCommitWithMessage(message, time string) {

	// message = message + ""
	// exec.Command("git", "commit", "-m", message)
	fmt.Println("git commit -m " + message)
}

func gitPush(automaticPush bool, branch string) {
	if automaticPush {
		//exec.Command("git", "push")
		fmt.Println("git push -u origin " + branch)
	}
}

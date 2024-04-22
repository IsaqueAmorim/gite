package command

import (
	"fmt"
	"os/exec"
	"slices"

	"github.com/IsaqueAmorim/gite/entity/branch"
)

func Commit(message string) {

	branch := branch.GetCurrentBranch()
	var protectedBranchs = []string{"release", "main", "master"}

	if !slices.Contains(protectedBranchs, branch) {

		exec.Command("git", "push", "-u", "origin", branch)
		return
	}

	fmt.Println("Você não pode fazer commit em um branch protegido!")
}

func GitAddAll(mustAdd bool) {
	if mustAdd {
		exec.Command("git", "add", "--all")
	}
}

func GitCommitWithMessage(message, time string) {

	message = message + ""
	exec.Command("git", "commit", "-m", message)
}

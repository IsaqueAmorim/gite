package branch

import (
	"fmt"
	"os/exec"
	"strings"

	stringHelper "github.com/IsaqueAmorim/gite/helpers"
)

type Branch struct {
	Prefix string
	Name   string
}

func GetCurrentBranch() string {
	cmd := exec.Command("git", "branch", "--show-current")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}

	return string(stdout)
}

func (b Branch) GetPrefix() string {
	return b.Prefix
}

func (b Branch) GetFullBranchName() string {
	return b.Prefix + "-" + b.Name
}

func (b Branch) GetName() string {
	return b.Name
}

func NewBranch(prefix, name string) *Branch {
	return &Branch{
		Prefix: prefix,
		Name:   getFormatedName(name),
	}
}

func getFormatedName(name string) string {
	str := stringHelper.Normalize(name)
	str = stringHelper.RemoveSpecialCharacters(str)
	str = stringHelper.AddSeparators(str)
	str = stringHelper.RemoveSequeneOfSeparators(str)

	return strings.ToLower(str)
}

// func StringToBranch(name string) *Branch {

// 	var prefix, branchName string

// 	if nameArr := strings.Split(name, "-"); len(nameArr) > 1 {
// 		prefix = nameArr[0]
// 		branchName = strings.Join(nameArr[1:], "-")
// 	} else {
// 		prefix = ""
// 		branchName = name
// 	}

// 	return &Branch{
// 		Prefix: prefix,
// 		Name:   branchName,
// 	}
// }

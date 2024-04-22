package branch

type Branch struct {
	Number int
	Title  string
}

func GetCurrentBranch() Branch {
	return Branch{
		Number: 1,
		Title:  "master",
	}
}

func (b Branch) GetNumber() int {
	return b.Number
}

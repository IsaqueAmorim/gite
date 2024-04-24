package commit

import (
	"slices"
	"strings"

	"github.com/IsaqueAmorim/gite/entity/branch"
)

type Commit struct {
	Message string
	Ticket  []string
	Time    string
	Status  string
}

func (t *Commit) SetTickets(tickets string) {

	ticketsArr := strings.Split(tickets, ",")

	if slices.Contains(ticketsArr, branch.GetCurrentBranch()) {
		t.Ticket = append(t.Ticket, ticketsArr...)
	} else {
		t.Ticket = append(t.Ticket, branch.GetCurrentBranch())
	}
}

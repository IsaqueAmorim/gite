package main

import (
	"fmt"

	"github.com/IsaqueAmorim/gite/entity/branch"
)

func main() {
	fmt.Println(branch.NewBranch("132020", "[TaxPlus] [Livro Declaração de Faturamento] - Valor do IPI na nota de Devolução não está subtraindo").GetFullBranchName())
}

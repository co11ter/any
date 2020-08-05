package main

import (
	"fmt"
	"github.com/co11ter/any/internal/coma"
	"github.com/co11ter/any/internal/comb"
	"github.com/co11ter/any/internal/facade"
)

func main() {
	fmt.Println(facade.NewFacade(coma.NewComponentA(), comb.NewComponentB()).Do())
}

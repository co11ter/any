package main

import (
	"fmt"
	"github.com/co11ter/any/internal/coma"
	"github.com/co11ter/any/internal/comb"
	"github.com/co11ter/any/internal/visitor"
)

func main() {
	a := coma.NewComponentA()
	b := comb.NewComponentB()

	fmt.Println(a.Accept(visitor.NewVisitor()))
	fmt.Println(b.Accept(visitor.NewVisitor()))
}

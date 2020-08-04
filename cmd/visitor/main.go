package main

import (
	"fmt"
	"github.com/co11ter/any/internal/visitor"
)

func main() {
	coma := visitor.NewComponentA()
	comb := visitor.NewComponentB()

	fmt.Println(coma.Accept(visitor.NewComVisitor()))
	fmt.Println(comb.Accept(visitor.NewComVisitor()))
}

package main

import (
	"fmt"
	"github.com/co11ter/any/internal/sub_a"
	"github.com/co11ter/any/internal/sub_b"
	"github.com/co11ter/any/internal/visitor"
)

func main() {
	coma := sub_a.NewComponentA()
	comb := sub_b.NewComponentB()

	fmt.Println(coma.Accept(visitor.NewComVisitor()))
	fmt.Println(comb.Accept(visitor.NewComVisitor()))
}

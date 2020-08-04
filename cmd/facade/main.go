package main

import (
	"fmt"
	"github.com/co11ter/any/internal/facade"
)

func main() {
	fmt.Println(facade.NewFacade().Do())
}

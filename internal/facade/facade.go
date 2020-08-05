package facade

import (
	"strings"
)

type componentA interface {
	// Op returns example
	Op() string
}

type componentB interface {
	// Op returns example
	Op() string
}

// Facade has to implement method for outside calling.
type Facade interface {
	Do() string
}

type facade struct {
	subA componentA
	subB componentB
}

// Do return result of component's processing.
func (f *facade) Do() string {
	result := []string{
		f.subA.Op(),
		f.subB.Op(),
	}

	return strings.Join(result, " ")
}

// Create new Facade.
func NewFacade(comA componentA, comB componentB) Facade {
	return &facade{
		subA: comA,
		subB: comB,
	}
}

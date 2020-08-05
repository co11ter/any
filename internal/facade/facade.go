package facade

import (
	"strings"
)

type component interface {
	Op() string
}

type Processor interface {
	Do() string
}

type facade struct {
	subA component
	subB component
}

func (f *facade) Do() string {
	result := []string{
		f.subA.Op(),
		f.subB.Op(),
	}
	return strings.Join(result, " ")
}

func NewFacade(comA, comB component) Processor {
	return &facade{
		subA: comA,
		subB: comB,
	}
}

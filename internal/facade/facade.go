package facade

import (
	"strings"
)

type Facade struct {
	subA *subsystemA
	subB *subsystemB
}

func NewFacade() *Facade {
	return &Facade{
		subA: newSubsystemA(),
		subB: newSubsystemB(),
	}
}

func (f *Facade) Do() string {
	result := []string{
		f.subA.Op(),
		f.subB.Op(),
	}
	return strings.Join(result, " ")
}

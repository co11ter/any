package any

import (
	"strings"
)

type Facade struct {
	subA *SubsystemA
	subB *SubsystemB
}

func NewFacade(subA *SubsystemA, subB *SubsystemB) *Facade {
	return &Facade{
		subA: subA,
		subB: subB,
	}
}

func (f *Facade) Do() string {
	result := []string{
		f.subA.Op(),
		f.subB.Op(),
	}
	return strings.Join(result, " ")
}

type SubsystemA struct{}

func (s *SubsystemA) Op() string { return "Hello" }

type SubsystemB struct{}

func (s *SubsystemB) Op() string { return "World" }

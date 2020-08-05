package comb

import (
	"github.com/co11ter/any/internal/visitor"
)

// ComponentA has to implements internal methods.
type ComponentB interface {
	// Accept processes visitor
	Accept(visitor.Visitor) string

	// Call returns example for visitor
	Call() string

	// Op returns example for facade
	Op() string
}

type componentB struct {
}

// Accept processes visitor.
func (c *componentB) Accept(v visitor.Visitor) string {
	return v.VisitB(c)
}

// Call returns example for visitor.
func (c *componentB) Call() string {
	return "Call B"
}

// Op returns example for facade.
func (c *componentB) Op() string {
	return "World"
}

// Create new ComponentB.
func NewComponentB() ComponentB {
	return &componentB{}
}

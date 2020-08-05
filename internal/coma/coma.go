package coma

import (
	"github.com/co11ter/any/internal/visitor"
)

// ComponentA has to implements internal methods.
type ComponentA interface {
	// Accept processes visitor
	Accept(visitor.Visitor) string

	// Call returns example for visitor
	Call() string

	// Op returns example for facade
	Op() string
}

type componentA struct {
}

// Accept processes visitor.
func (c *componentA) Accept(v visitor.Visitor) string {
	return v.VisitA(c)
}

// Call returns example for visitor.
func (c *componentA) Call() string {
	return "Call A"
}

// Op returns example for facade.
func (c *componentA) Op() string {
	return "Hello"
}

// Create new ComponentA.
func NewComponentA() ComponentA {
	return &componentA{}
}

package coma

import (
	"github.com/co11ter/any/internal/visitor"
)

type ComponentA interface {
	Accept(visitor.Visitor) string
	Call() string
	Op() string
}

type componentA struct {
}

func (c *componentA) Accept(v visitor.Visitor) string {
	return v.VisitA(c)
}

func (c *componentA) Call() string {
	return "Call A"
}

func (c *componentA) Op() string {
	return "Hello"
}

func NewComponentA() ComponentA {
	return &componentA{}
}

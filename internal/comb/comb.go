package comb

import (
	"github.com/co11ter/any/internal/visitor"
)

type ComponentB interface {
	Accept(visitor.Visitor) string
	Call() string
	Op() string
}

type componentB struct {
}

func (c *componentB) Accept(v visitor.Visitor) string {
	return v.VisitB(c)
}

func (c *componentB) Call() string {
	return "Call B"
}

func (c *componentB) Op() string {
	return "World"
}

func NewComponentB() ComponentB {
	return &componentB{}
}

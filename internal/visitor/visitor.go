package visitor

import (
	"github.com/co11ter/any/internal/sub_a"
	"github.com/co11ter/any/internal/sub_b"
)

type Visitor interface {
	VisitA(*sub_a.ComponentA) string
	VisitB(*sub_b.ComponentB) string
}

type acceptor interface {
	Accept(v Visitor) string
}

type ComVisitor struct {
}

func NewComVisitor() *ComVisitor {
	return &ComVisitor{}
}

func (v *ComVisitor) VisitA(comA *sub_a.ComponentA) string {
	return "Visit: " + comA.Call()
}

func (v *ComVisitor) VisitB(comB *sub_b.ComponentB) string {
	return "Visit: " + comB.Call()
}

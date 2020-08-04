package visitor

type Visitor interface {
	VisitA() string
	VisitB() string
}

type Acceptor interface {
	Accept(v Visitor) string
}

type ComponentsVisitor struct {
}

func (v *ComponentsVisitor) VisitA(comA *ComponentA) string {
	return comA.Call()
}

func (v *ComponentsVisitor) VisitB(comB *ComponentB) string {
	return comB.Call()
}

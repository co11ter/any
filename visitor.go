package any

type Visitor interface {
	VisitA() string
	VisitB() string
}

type Acceptor interface {
	Accept(v Visitor) string
}

type ComponentA struct{}

func (c *ComponentA) Accept(v Visitor) string {
	return v.VisitA()
}

func (c *ComponentA) Call() string {
	return "Call A"
}

type ComponentB struct{}

func (c *ComponentB) Accept(v Visitor) string {
	return v.VisitB()
}

func (c *ComponentB) Call() string {
	return "Call B"
}

type ComponentsVisitor struct{}

func (v *ComponentsVisitor) VisitA(comA *ComponentA) string {
	return comA.Call()
}

func (v *ComponentsVisitor) VisitB(comB *ComponentB) string {
	return comB.Call()
}

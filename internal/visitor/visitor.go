package visitor

type visitor interface {
	VisitA(*ComponentA) string
	VisitB(*ComponentB) string
}

type acceptor interface {
	Accept(v visitor) string
}

type ComVisitor struct {
}

func NewComVisitor() *ComVisitor {
	return &ComVisitor{}
}

func (v *ComVisitor) VisitA(comA *ComponentA) string {
	return "Visit: " + comA.Call()
}

func (v *ComVisitor) VisitB(comB *ComponentB) string {
	return "Visit: " + comB.Call()
}

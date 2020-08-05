package visitor

type caller interface {
	Call() string
}

type Visitor interface {
	VisitA(caller) string
	VisitB(caller) string
}

type visitor struct {
}

func (v *visitor) VisitA(comA caller) string {
	return "Visit: " + comA.Call()
}

func (v *visitor) VisitB(comB caller) string {
	return "Visit: " + comB.Call()
}

func NewVisitor() Visitor {
	return &visitor{}
}

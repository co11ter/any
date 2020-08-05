package visitor

type Caller interface {
	Call() string
}

type Visitor interface {
	VisitA(Caller) string
	VisitB(Caller) string
}

type visitor struct {
}

func (v *visitor) VisitA(comA Caller) string {
	return "Visit: " + comA.Call()
}

func (v *visitor) VisitB(comB Caller) string {
	return "Visit: " + comB.Call()
}

func NewVisitor() Visitor {
	return &visitor{}
}

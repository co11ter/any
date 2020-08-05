package visitor

type componentA interface {
	// Call returns example
	Call() string
}

type componentB interface {
	// Call returns example
	Call() string
}

// Visitor has to implement methods for visiting components.
type Visitor interface {
	// VisitA visits components A
	VisitA(componentA) string

	// VisitB visits components B
	VisitB(componentB) string
}

type visitor struct {
}

// VisitA visits components A.
func (v *visitor) VisitA(comA componentA) string {
	return "Visit: " + comA.Call()
}

// VisitB visits components B.
func (v *visitor) VisitB(comB componentB) string {
	return "Visit: " + comB.Call()
}

// Create new visitor.
func NewVisitor() Visitor {
	return &visitor{}
}

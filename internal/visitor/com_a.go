package visitor

type ComponentA struct {
}

func (c *ComponentA) Accept(v Visitor) string {
	return v.VisitA()
}

func (c *ComponentA) Call() string {
	return "Call A"
}

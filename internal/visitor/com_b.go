package visitor

type ComponentB struct {
}

func (c *ComponentB) Accept(v Visitor) string {
	return v.VisitB()
}

func (c *ComponentB) Call() string {
	return "Call B"
}

package visitor

type ComponentB struct {
}

func NewComponentB() *ComponentB {
	return &ComponentB{}
}

func (c *ComponentB) Accept(v visitor) string {
	return v.VisitB(c)
}

func (c *ComponentB) Call() string {
	return "Call B"
}

package sub_a

type visitor interface {
	VisitA(*ComponentA) string
}

type ComponentA struct {
}

func NewComponentA() *ComponentA {
	return &ComponentA{}
}

func (c *ComponentA) Accept(v visitor) string {
	return v.VisitA(c)
}

func (c *ComponentA) Call() string {
	return "Call A"
}

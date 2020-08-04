package facade

type subsystemB struct {
}

func newSubsystemB() *subsystemB {
	return &subsystemB{}
}

func (s *subsystemB) Op() string {
	return "World"
}

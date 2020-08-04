package facade

type subsystemA struct {
}

func newSubsystemA() *subsystemA {
	return &subsystemA{}
}

func (s *subsystemA) Op() string {
	return "Hello"
}

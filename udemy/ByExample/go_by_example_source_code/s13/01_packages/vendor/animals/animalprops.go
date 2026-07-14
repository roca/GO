package animals

// Speaker ... your comments
type Speaker interface {
	Speaks() string
}

// Perform ... your comments
func Perform(s Speaker) string { return s.Speaks() }

// Lion ... your comments
type Lion struct{}

// Speaks ... your comments
func (s Lion) Speaks() string { return "roars" }

// Cat ... your comments
type Cat struct{}

// Speaks ... your comments
func (s Cat) Speaks() string { return "meaws" }

// Human ... your comments
type Human struct{}

// Speaks ... your comments
func (s Human) Speaks() string { return "talks" }

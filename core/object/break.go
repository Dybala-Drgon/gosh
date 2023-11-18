package object

// Break represents a break statement.
type Break struct{}

// TypeName returns the name of the type.
func (o *Break) TypeName() string {
	return "break"
}

func (o *Break) String() string {
	return "<break>"
}

// IsFalsy returns true if the value of the type is falsy.
func (o *Break) IsFalsy() bool {
	return false
}

// Equals returns true if the value of the type
// is equal to the value of another object.
func (o *Break) Equals(x Object) bool {
	return false
}

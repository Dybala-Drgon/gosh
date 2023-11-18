package object

// Function represents a user function.
type Function struct {
	Value CallableFunc
}

// TypeName returns the name of the type.
func (o *Function) TypeName() string {
	return "user-function"
}

func (o *Function) String() string {
	return "<user-function>"
}

// IsFalsy returns true if the value of the type is falsy.
func (o *Function) IsFalsy() bool {
	return false
}

// Equals returns true if the value of the type
// is equal to the value of another object.
func (o *Function) Equals(x Object) bool {
	return false
}

// Call invokes a user function.
func (o *Function) Call(args ...Object) ([]Object, error) {
	return o.Value(args...)
}

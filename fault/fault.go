// Package fault holds the Fault structure and its functionality
package fault

type (
	// Type of the Fault (Error, Warning, Info or Debug)
	Type int

	// Fault struct represents an Error, Warning, Info or Debug fault
	Fault struct {
		typ     Type
		message string
	}
)

// Error returns the message string of the Fault
func (f *Fault) Error() string {
	return f.message
}

// String returns the Fault message as string
func (f *Fault) String() string {
	return f.message
}

// Is check the Fault for a Type
func (f *Fault) Is(typ Type) bool {
	switch typ {
	case Error:
		return f.typ == Error
	case Warning:
		return f.typ == Warning
	case Info:
		return f.typ == Info
	case Debug:
		return f.typ == Debug
	}
	return false
}

// OneOf check the Fault for many Type
func (f *Fault) OneOf(typ ...Type) bool {
	for _, t := range typ {
		if f.Is(t) {
			return true
		}
	}
	return false
}

const (
	// Error fault Type
	Error Type = iota
	// Warning fault Type
	Warning
	// Info fault Type
	Info
	// Debug fault Type
	Debug
)

// NewError create a new Error Object
func NewError(msg string) *Fault {
	return &Fault{
		typ:     Error,
		message: msg,
	}
}

// NewWarning create a new Warning Object
func NewWarning(msg string) *Fault {
	return &Fault{
		typ:     Warning,
		message: msg,
	}
}

// NewInfo create a new Info Object
func NewInfo(msg string) *Fault {
	return &Fault{
		typ:     Info,
		message: msg,
	}
}

// NewDebug create a new Debug Object
func NewDebug(msg string) *Fault {
	return &Fault{
		typ:     Debug,
		message: msg,
	}
}

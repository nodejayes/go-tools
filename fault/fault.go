package fault

type (
	FaultType int

	Fault struct {
		typ     FaultType
		message string
	}
)

func (f *Fault) Error() string {
	return f.message
}

func (f *Fault) Is(typ FaultType) bool {
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

func (f *Fault) OneOf(typ ...FaultType) bool {
	for _, t := range typ {
		if f.Is(t) {
			return true
		}
	}
	return false
}

const (
	Error FaultType = iota
	Warning
	Info
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

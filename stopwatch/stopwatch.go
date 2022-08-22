package stopwatch

const (
	// ToMillisecondsDivisor a divisor to convert ns into ms
	ToMillisecondsDivisor = 1000000.0
)

type (
	Stopwatch struct {
		start    int64
		sessions map[string]int64
	}
)

func New() *Stopwatch {
	return &Stopwatch{
		sessions: make(map[string]int64),
	}
}

package counter

type (
	Counter struct {
		value int
	}
)

// New create a new Counter begins by 0 on default
func New(initialValue ...int) *Counter {
	if len(initialValue) < 1 {
		return &Counter{value: 0}
	}
	return &Counter{value: initialValue[0]}
}

// GetValue returns the current Counter value
func (c *Counter) GetValue() int {
	return c.value
}

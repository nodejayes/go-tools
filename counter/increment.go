package counter

// Increment the Counter
func (c *Counter) Increment(step ...int) {
	if len(step) < 1 {
		c.value++
		return
	}
	c.value += step[0]
}

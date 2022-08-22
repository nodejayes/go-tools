package counter

// Decrement the Counter
func (c *Counter) Decrement(step ...int) {
	if len(step) < 1 {
		if c.value-1 < 0 {
			c.value = 0
			return
		}
		c.value--
		return
	}
	if c.value-step[0] < 0 {
		c.value = 0
		return
	}
	c.value -= step[0]
}

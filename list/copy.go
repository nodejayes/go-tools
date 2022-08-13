package list

// Copy the current List into a new List
func (l List[T]) Copy() List[T] {
	var c List[T]
	c.innerList = l.innerList[:len(l.innerList)]
	return c
}

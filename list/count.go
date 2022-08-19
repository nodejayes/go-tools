package list

// Count get the number of Elements in the List
func (l *List[T]) Count() int {
	return len(l.innerList)
}

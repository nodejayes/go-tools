package list

// Clear removes all Items in the List
func (l *List[T]) Clear() {
	l.innerList = make([]T, 0)
}

package list

// Find finds a Item in a List of Items and returns it
func (l List[T]) Find(predicate func(T, int, List[T]) bool) T {
	for idx, item := range l.innerList {
		if predicate(item, idx, l) {
			return item
		}
	}
	return *new(T)
}

package list

// Find find a Item in a List of Items and returns it
func (l List[T]) Find(predicate func(T, int, []T) bool) T {
	for idx, item := range l.innerList {
		if predicate(item, idx, l.innerList) {
			return item
		}
	}
	return *new(T)
}

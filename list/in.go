package list

// In checks if a Item in a List of Items
func (l *List[T]) In(item T, predicate func(T, T) bool) bool {
	for idx := range l.innerList {
		if predicate(l.innerList[idx], item) {
			return true
		}
	}
	return false
}

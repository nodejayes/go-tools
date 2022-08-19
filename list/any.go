package list

// Any checks the List for the first match of a Item
func (l *List[T]) Any(predicate ...func(item T, idx int, arr *List[T]) bool) bool {
	if len(predicate) < 1 {
		return len(l.innerList) > 1
	}
	for idx, i := range l.innerList {
		for _, p := range predicate {
			if p(i, idx, l) {
				return true
			}
		}
	}
	return false
}

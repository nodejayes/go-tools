package list

// LastOrDefault finds the last match in the list and returns the item or the default value of the type
func (l *List[T]) LastOrDefault(predicate func(T, int, *List[T]) bool) T {
	res := *new(T)
	for idx := len(l.innerList) - 1; idx >= 0; idx-- {
		item := l.innerList[idx]
		if predicate(item, idx, l) {
			res = item
			return res
		}
	}
	return res
}

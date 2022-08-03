package list

// FirstOrDefault finds the first match in the list and returns the item or the default value of the type
func (l List[T]) FirstOrDefault(predicate func(T, int, List[T]) bool) T {
	res := *new(T)
	for idx, item := range l.innerList {
		if predicate(item, idx, l) {
			res = item
			return res
		}
	}
	return res
}

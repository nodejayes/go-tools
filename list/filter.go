package list

// Filter removes all Items from a List that not Equals the predicate Function
func (l *List[T]) Filter(predicate func(T, int, *List[T]) bool) *List[T] {
	result := New(make([]T, 0))

	for idx, item := range l.innerList {
		if predicate(item, idx, l) {
			result.innerList = append(result.innerList, item)
		}
	}

	return result
}

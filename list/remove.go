package list

// Remove a Item from List
func (l *List[T]) Remove(predicate func(item T, idx int, arr *List[T]) bool) {
	for idx, i := range l.innerList {
		if predicate(i, idx, l) {
			l.innerList = append(l.innerList[:idx], l.innerList[idx+1:]...)
		}
	}
}

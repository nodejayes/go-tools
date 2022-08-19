package list

// Pull a Item from the List and give it back
func (l *List[T]) Pull(idx int) T {
	if idx > l.Count()-1 || idx < 0 {
		return *new(T)
	}
	item := l.innerList[idx]
	l.Remove(func(_ T, index int, _ *List[T]) bool {
		return index == idx
	})
	return item
}

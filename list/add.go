package list

// Add a new Item to the List
func (l *List[T]) Add(item T) {
	l.innerList = append(l.innerList, item)
}

// AddIfNotExists add a Item to the List when not exists on the list
func (l *List[T]) AddIfNotExists(item T, predicate ...func(T, int, *List[T]) bool) {
	if len(predicate) < 1 {
		predicate = append(predicate, func(t T, i int, l *List[T]) bool {
			return interface{}(t) == interface{}(item)
		})
	}
	if l.Any(predicate...) {
		return
	}
	l.Add(item)
}

// AddRange adds multiple Items to the List
func (l *List[T]) AddRange(items *List[T]) {
	for _, item := range items.GetItems() {
		l.Add(item)
	}
}

// AddRangeIfNotExists add multiple Items to the List when not exists
func (l *List[T]) AddRangeIfNotExists(items *List[T], predicate ...func(T, int, *List[T]) bool) {
	for _, item := range items.GetItems() {
		l.AddIfNotExists(item, predicate...)
	}
}

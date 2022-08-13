package list

// Map converts a List of Items into a other List of other Items
func Map[T, K any](items List[T], transform func(T, int, List[T]) K) List[K] {
	result := New(make([]K, len(items.innerList)))

	for idx, item := range items.innerList {
		result.innerList[idx] = transform(item, idx, items)
	}

	return result
}

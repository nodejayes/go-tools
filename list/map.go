package list

// Map converts a List of Items into a other List of other Items
func Map[T, K any](items List[T], transform func(T, int, List[T]) K) []K {
	result := make([]K, len(items.innerList))

	for idx, item := range items.innerList {
		result[idx] = transform(item, idx, items)
	}

	return result
}

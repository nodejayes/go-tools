package list

// Map converts a List of Items into a other List of other Items
func Map[T, K any](items []T, transform func(T, int, []T) K) []K {
	result := make([]K, len(items))

	for idx, item := range items {
		result[idx] = transform(item, idx, items)
	}

	return result
}

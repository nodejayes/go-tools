package list

// Reduce decrease a List of Items to a complete new Object
func Reduce[T, K any](items List[T], transform func(K, T, List[T]) K, initial K) K {
	for _, item := range items.innerList {
		initial = transform(initial, item, items)
	}

	return initial
}

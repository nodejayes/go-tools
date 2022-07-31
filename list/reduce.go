package list

// Reduce decrease a List of Items to a complete new Object
func Reduce[T, K any](items []T, transform func(K, T, []T) K, initial K) K {
	for _, item := range items {
		initial = transform(initial, item, items)
	}

	return initial
}

// Package list contains helper Functions for Go Slices
package list

type (
	List[T any] struct {
		innerList []T
	}
)

// New creates a new Instance of a List
func New[T any](items []T) *List[T] {
	return &List[T]{
		innerList: items,
	}
}

// GetItems return a Slice of the current Items in the List
func (l *List[T]) GetItems() []T {
	return l.innerList
}

// SetItems set the current Items in the List to the given Slice od Items
func (l *List[T]) SetItems(items []T) {
	l.innerList = items
}

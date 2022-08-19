package list

import "fmt"

// FindIndex get the Index of a Item in the List
func (l *List[T]) FindIndex(predicate func(T, int, *List[T]) bool) (int, error) {
	for idx, item := range l.innerList {
		if predicate(item, idx, l) {
			return idx, nil
		}
	}
	return 0, fmt.Errorf("item not found in list")
}

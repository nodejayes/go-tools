package list

import "fmt"

// ElementAt finds the item at the index number and returns it
// when the list has not enough items a error was returned
func (l *List[T]) ElementAt(idx int) (T, error) {
	if idx > len(l.innerList)-1 {
		return *new(T), fmt.Errorf("element with index %v not in list. list has only %v elements", idx, len(l.innerList))
	}
	return l.innerList[idx], nil
}

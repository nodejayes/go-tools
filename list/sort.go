package list

import "sort"

// Sort
func (l List[T]) Sort(sortable func(T, T) bool) {
	sort.Slice(l.innerList, func(i, j int) bool {
		return sortable(l.innerList[i], l.innerList[j])
	})
}

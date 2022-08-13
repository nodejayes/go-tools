package list

// ForSegment iterates over a segment in a List and do something wit the item before and current item
func (l List[T]) ForSegment(toDo func(T, T)) {
	for idx := range l.innerList {
		if idx == 0 {
			continue
		}
		toDo(l.innerList[idx-1], l.innerList[idx])
	}
}

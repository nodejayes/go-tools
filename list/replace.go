package list

// Replace a Item in the List at the index
func (l *List[T]) Replace(item T, idx int) {
	if idx > l.Count()-1 || idx < 0 {
		return
	}
	l.innerList[idx] = item
}

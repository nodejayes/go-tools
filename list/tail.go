package list

// Tail get items from the List with the given length from the end
func (l List[T]) Tail(length int) List[T] {
	var res List[T]
	if length < 0 {
		length = 0
	}
	if length > l.Count() {
		return l.Copy()
	}
	res.innerList = l.innerList[l.Count()-length:]
	return res
}

package list

// Head get items from the List with the given length from the begin
func (l *List[T]) Head(length int) *List[T] {
	var res List[T]
	if length < 0 {
		length = 0
	}
	if length > l.Count() {
		return l.Copy()
	}
	res.innerList = l.innerList[:length]
	return &res
}

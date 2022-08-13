package list

// Reverse turns around the current List
func (l List[T]) Reverse() List[T] {
	a := l.Copy()
	for i := a.Count()/2 - 1; i >= 0; i-- {
		opp := a.Count() - 1 - i
		a.innerList[i], a.innerList[opp] = a.innerList[opp], a.innerList[i]
	}
	return a
}

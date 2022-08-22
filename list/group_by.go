package list

// GroupBy converts the items in the List to a string that represents the group key and returns a Dictionary of new List group by the group key
func (l *List[T]) GroupBy(converter func(T) string) map[string]*List[T] {
	res := map[string]*List[T]{}
	for _, item := range l.GetItems() {
		key := converter(item)
		if res[key] == nil {
			res[key] = New(make([]T, 0))
		}
		res[key].Add(item)
	}
	return res
}

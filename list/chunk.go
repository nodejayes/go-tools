package list

// Chunk split the List into smaller Lists
func (l *List[T]) Chunk(chunkSize int) []*List[T] {
	var res []*List[T]
	if chunkSize < 1 {
		res = append(res, l)
		return res
	}

	var chunk List[T]
	for _, item := range l.innerList {
		chunk.Add(item)
		if chunk.Count() >= chunkSize {
			res = append(res, chunk.Copy())
			chunk.Clear()
		}
	}
	if chunk.Count() > 0 {
		res = append(res, chunk.Copy())
		chunk.Clear()
	}
	return res
}

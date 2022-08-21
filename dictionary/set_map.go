package dictionary

// SetMap set the internal Key Value Pairs of the Dictionary
func (d *Dictionary[T, K]) SetMap(data map[T]K) {
	if data == nil {
		data = make(map[T]K)
	}
	d.data = data
}

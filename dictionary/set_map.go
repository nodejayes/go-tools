package dictionary

// SetMap set the internal Map of the Dictionary
func (d *Dictionary[T, K]) SetMap(data map[T]K) {
	d.data = data
}

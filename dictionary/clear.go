package dictionary

// Clear the internal Key Value Pairs
func (d *Dictionary[T, K]) Clear() {
	d.data = make(map[T]K)
}

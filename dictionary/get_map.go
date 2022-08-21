package dictionary

// GetMap get the internal Key Value Pairs of the Dictionary as Map
func (d *Dictionary[T, K]) GetMap() map[T]K {
	return d.data
}

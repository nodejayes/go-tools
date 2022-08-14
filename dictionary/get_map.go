package dictionary

// GetMap get the internal Map of the Dictionary
func (d Dictionary[T, K]) GetMap() map[T]K {
	return d.data
}

package dictionary

// Add a new Key Value Pair to the Dictionary
func (d *Dictionary[T, K]) Add(key T, value K) {
	d.data[key] = value
}

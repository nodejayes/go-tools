package dictionary

// Count the Entries in the Dictionary
func (d *Dictionary[T, K]) Count() int {
	return len(d.data)
}

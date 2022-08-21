package dictionary

// Count the Entries in the Dictionary and get the number of Entries
func (d *Dictionary[T, K]) Count() int {
	return len(d.data)
}

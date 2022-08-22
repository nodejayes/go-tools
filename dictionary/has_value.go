package dictionary

// HasValue checks the Dictionary for a value to exists
func (d *Dictionary[T, K]) HasValue(value K) bool {
	found := false
	for _, k := range d.Values().GetItems() {
		if interface{}(k) == interface{}(value) {
			found = true
			break
		}
	}
	return found
}

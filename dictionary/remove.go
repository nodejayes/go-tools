package dictionary

// Remove a Value from Dictionary
func (d *Dictionary[T, K]) Remove(key T) {
	found := false
	for _, k := range d.Keys().GetItems() {
		if k == key {
			found = true
			break
		}
	}

	if !found {
		return
	}
	delete(d.data, key)
}

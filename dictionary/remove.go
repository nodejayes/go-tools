package dictionary

import "github.com/nodejayes/go-tools/v3/compare"

// Remove a Value from Dictionary
func (d *Dictionary[T, K]) Remove(key T) {
	if !d.Keys().Any(compare.FilterEquals(key)) {
		return
	}
	delete(d.data, key)
}

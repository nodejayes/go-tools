package dictionary

import (
	"fmt"
	"github.com/nodejayes/go-tools/v3/compare"
	"github.com/nodejayes/go-tools/v3/fault"
	"github.com/nodejayes/go-tools/v3/list"
)

// Values returns the Values of the Dictionary
func (d *Dictionary[T, K]) Values() *list.List[K] {
	values := list.New(make([]K, 0))
	for _, val := range d.data {
		values.Add(val)
	}
	return values
}

// GetValue returns the Dictionary Value by the given key. if no Entry with the key in the Dictionary a Fault was returned.
func (d *Dictionary[T, K]) GetValue(key T) (K, *fault.Fault) {
	if !d.Keys().Any(compare.FilterEquals(key)) {
		return *new(K), fault.NewError(fmt.Sprintf("key not found in Dictionary %v", key))
	}
	return d.data[key], nil
}

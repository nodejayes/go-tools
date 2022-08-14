package dictionary

import "github.com/nodejayes/go-tools/list"

// Values returns the Values of the Dictionary
func (d Dictionary[T, K]) Values() list.List[K] {
	values := list.New(make([]K, 0))
	for _, val := range d.data {
		values.Add(val)
	}
	return values
}

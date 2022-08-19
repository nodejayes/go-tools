package dictionary

import "github.com/nodejayes/go-tools/v3/list"

// Keys returns the Keys of the Dictionary
func (d *Dictionary[T, K]) Keys() *list.List[T] {
	keyList := list.New(make([]T, 0))
	for key := range d.data {
		keyList.Add(key)
	}
	return keyList
}

// Package dictionary defines the Dictionary and its Functionality
package dictionary

import (
	"bytes"
	"fmt"
)

type (
	// Dictionary represents a List of Key Value Pairs
	Dictionary[T comparable, K any] struct {
		data map[T]K
	}
)

// String returns the Dictionary as string
func (d *Dictionary[T, K]) String() string {
	buf := bytes.NewBuffer([]byte{})
	keys := d.Keys()
	for _, key := range keys.GetItems() {
		if buf.Len() > 0 {
			buf.WriteString(",")
		}
		value, f := d.GetValue(key)
		if f != nil {
			continue
		}
		buf.WriteString(fmt.Sprintf("%v:%v", key, value))
	}
	return fmt.Sprintf("{%v}", buf.String())
}

// New create a Dictionary
func New[T comparable, K any](data map[T]K) *Dictionary[T, K] {
	if data == nil {
		data = make(map[T]K)
	}
	return &Dictionary[T, K]{
		data,
	}
}

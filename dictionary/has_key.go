package dictionary

import (
	"github.com/nodejayes/go-tools/v3/list"
)

// HasKey check the Dictionary for a specific Key
func (d *Dictionary[T, K]) HasKey(key T) bool {
	return d.Keys().Any(func(item T, _ int, _ *list.List[T]) bool {
		return interface{}(item) == interface{}(key)
	})
}

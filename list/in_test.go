package list_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/compare"
	"github.com/nodejayes/go-tools/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIn(t *testing.T) {
	t.Run("find item in list", func(t *testing.T) {
		assert.True(t, list.New([]int{1, 2, 3}).In(2, compare.Equals[int]))
	})

	t.Run("not find item in list", func(t *testing.T) {
		assert.False(t, list.New([]int{1, 2, 3}).In(5, compare.Equals[int]))
	})

	t.Run("not find item in empty list", func(t *testing.T) {
		assert.False(t, list.New([]int{}).In(5, compare.Equals[int]))
	})
}

func ExampleList_In() {
	myList := list.New([]int{1, 2, 3})
	fmt.Println(myList.In(2, compare.Equals[int]))
	// Output: true
}

func BenchmarkIn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list.New([]int{1, 2, 3}).In(2, compare.Equals[int])
	}
}

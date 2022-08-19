package list_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/v3/compare"
	"github.com/nodejayes/go-tools/v3/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilter(t *testing.T) {
	t.Run("remove 2 from List [1 2 3]", func(t *testing.T) {
		assert.Equal(t, []int{1, 3}, list.New([]int{1, 2, 3}).Filter(compare.FilterNotEquals(2)).GetItems())
	})
}

func ExampleList_Filter() {
	myList := list.New([]int{1, 2, 3})
	fmt.Println(myList.Filter(compare.FilterNotEquals(2)).GetItems())
	// Output: [1 3]
}

func BenchmarkFilter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list.New([]int{1, 2, 3}).Filter(compare.FilterNotEquals(2))
	}
}

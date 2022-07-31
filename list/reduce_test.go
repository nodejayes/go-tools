package list_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReduce(t *testing.T) {
	t.Run("int list to sum", func(t *testing.T) {
		sum := list.Reduce([]int{1, 1, 1}, func(result int, item int, itemList []int) int {
			return item + result
		}, 0)
		assert.Equal(t, 3, sum)
	})
}

func ExampleReduce() {
	fmt.Println(list.Reduce([]int{1, 1, 1}, func(result int, item int, itemList []int) int {
		return item + result
	}, 0))
	// Output: 3
}

func BenchmarkReduce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list.Reduce([]int{1, 1, 1}, func(result int, item int, itemList []int) int {
			return item + result
		}, 0)
	}
}

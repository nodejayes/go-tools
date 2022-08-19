package list_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/v3/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReduce(t *testing.T) {
	t.Run("int list to sum", func(t *testing.T) {
		myList := list.New([]int{1, 1, 1})
		sum := list.Reduce(myList, func(result int, item int, itemList *list.List[int]) int {
			return item + result
		}, 0)
		assert.Equal(t, 3, sum)
	})
}

func ExampleReduce() {
	myList := list.New([]int{1, 1, 1})
	fmt.Println(list.Reduce(myList, func(result int, item int, itemList *list.List[int]) int {
		return item + result
	}, 0))
	// Output: 3
}

func BenchmarkReduce(b *testing.B) {
	myList := list.New([]int{1, 1, 1})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Reduce(myList, func(result int, item int, itemList *list.List[int]) int {
			return item + result
		}, 0)
	}
}

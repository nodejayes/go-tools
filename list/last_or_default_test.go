package list_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/compare"
	"github.com/nodejayes/go-tools/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_LastOrDefault(t *testing.T) {
	t.Run("find 5 in [1 2 3 4 5]", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3, 4, 5})
		assert.Equal(t, 5, myList.LastOrDefault(compare.FilterEquals(5)))
	})
	t.Run("find 1 in [1 2 3 4 5]", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3, 4, 5})
		assert.Equal(t, 1, myList.LastOrDefault(compare.FilterEquals(1)))
	})
	t.Run("find 3 in [1 2 3 4 5]", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3, 4, 5})
		assert.Equal(t, 3, myList.LastOrDefault(compare.FilterEquals(3)))
	})
	t.Run("get default when search 5 in [1 2 3]", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3})
		assert.Equal(t, 0, myList.LastOrDefault(compare.FilterEquals(5)))
	})
}

func ExampleList_LastOrDefault() {
	myList := list.New([]int{1, 2, 3, 4, 5})
	fmt.Println(myList.LastOrDefault(compare.FilterEquals(4)))
	// Output: 4
}

func BenchmarkList_LastOrDefault(b *testing.B) {
	myList := list.New([]int{1, 2, 3, 4, 5})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myList.LastOrDefault(compare.FilterEquals(4))
	}
}

package list_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/v3/compare"
	"github.com/nodejayes/go-tools/v3/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_FirstOrDefault(t *testing.T) {
	myList := list.New([]int{1, 2, 3})
	t.Run("find the first element from list", func(t *testing.T) {
		assert.Equal(t, 1, myList.FirstOrDefault(compare.FilterTrue[int]()))
	})

	t.Run("find 2 in list [1 2 3]", func(t *testing.T) {
		assert.Equal(t, 2, myList.FirstOrDefault(compare.FilterEquals(2)))
	})

	t.Run("empty list returns default value", func(t *testing.T) {
		emptyList := list.New([]int{})
		assert.Equal(t, 0, emptyList.FirstOrDefault(compare.FilterTrue[int]()))
	})
}

func ExampleList_FirstOrDefault() {
	myList := list.New([]int{1, 2, 3})
	el := myList.FirstOrDefault(compare.FilterEquals(2))
	fmt.Println(el)
	// Output: 2
}

func BenchmarkList_FirstOrDefault(b *testing.B) {
	myList := list.New([]int{1, 2, 3})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myList.FirstOrDefault(compare.FilterEquals(2))
	}
}

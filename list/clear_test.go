package list_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/v2/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_Clear(t *testing.T) {
	t.Run("clear a empty list", func(t *testing.T) {
		myList := list.New(make([]int, 0))
		myList.Clear()
		assert.Equal(t, []int{}, myList.GetItems())
	})
	t.Run("clear [1 2 3]", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3})
		myList.Clear()
		assert.Equal(t, []int{}, myList.GetItems())
	})
}

func ExampleList_Clear() {
	myList := list.New([]int{1, 2, 3})
	myList.Clear()
	fmt.Println(myList.GetItems())
	// Output: []
}

func BenchmarkList_Clear(b *testing.B) {
	myList := list.New([]int{1, 2, 3})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myList.Clear()
	}
}

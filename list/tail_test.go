package list_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/v2/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_Tail(t *testing.T) {
	t.Run("get the last 2 items from [1 2 3 4 5]", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3, 4, 5})
		assert.Equal(t, []int{4, 5}, myList.Tail(2).GetItems())
	})
	t.Run("get the last 2 items from [1]", func(t *testing.T) {
		myList := list.New([]int{1})
		assert.Equal(t, []int{1}, myList.Tail(2).GetItems())
	})
	t.Run("get the last 0 items from [1 2 3 4 5]", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3, 4, 5})
		assert.Equal(t, []int{}, myList.Tail(0).GetItems())
	})
	t.Run("get the last 0 items from [1 2 3 4 5]", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3, 4, 5})
		assert.Equal(t, []int{}, myList.Tail(-1).GetItems())
	})
}

func ExampleList_Tail() {
	myList := list.New([]int{1, 2, 3, 4, 5})
	fmt.Println(myList.Tail(2).GetItems())
	// Output: [4 5]
}

func BenchmarkList_Tail(b *testing.B) {
	myList := list.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myList.Tail(2)
	}
}

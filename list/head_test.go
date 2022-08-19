package list_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/v3/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_Head(t *testing.T) {
	t.Run("can get the first 2 from [1 2 3 4 5]", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3, 4, 5})
		assert.Equal(t, []int{1, 2}, myList.Head(2).GetItems())
	})
	t.Run("can get the first 2 from [1]", func(t *testing.T) {
		myList := list.New([]int{1})
		assert.Equal(t, []int{1}, myList.Head(2).GetItems())
	})
	t.Run("can get the first 0 from [1 2 3 4 5]", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3, 4, 5})
		assert.Equal(t, []int{}, myList.Head(0).GetItems())
	})
	t.Run("can get the first 0 from [1 2 3 4 5]", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3, 4, 5})
		assert.Equal(t, []int{}, myList.Head(-1).GetItems())
	})
}

func ExampleList_Head() {
	myList := list.New([]int{1, 2, 3, 4, 5})
	fmt.Println(myList.Head(2).GetItems())
	// Output: [1 2]
}

func BenchmarkList_Head(b *testing.B) {
	myList := list.New([]int{1, 2, 3, 4, 5})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myList.Head(2)
	}
}

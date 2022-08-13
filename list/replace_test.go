package list_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_Replace(t *testing.T) {
	t.Run("replace 2 with 5 in [1 2 3]", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3})
		myList.Replace(5, 1)
		assert.Equal(t, []int{1, 5, 3}, myList.GetItems())
	})
	t.Run("not replace 2 with 5 in [1 2 3] when idx < 0", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3})
		myList.Replace(5, -1)
		assert.Equal(t, []int{1, 2, 3}, myList.GetItems())
	})
	t.Run("not replace 2 with 5 in [1 2 3] when idx > list length", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3})
		myList.Replace(5, 3)
		assert.Equal(t, []int{1, 2, 3}, myList.GetItems())
	})
}

func ExampleList_Replace() {
	myList := list.New([]int{1, 2, 3})
	myList.Replace(5, 1)
	fmt.Println(myList.GetItems())
	// Output: [1 5 3]
}

func BenchmarkList_Replace(b *testing.B) {
	myList := list.New([]int{1, 2, 3})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myList.Replace(5, 1)
	}
}

package list_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/v2/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList(t *testing.T) {
	t.Run("can get inner Items", func(t *testing.T) {
		assert.Equal(t, []int{1, 2, 3}, list.New([]int{1, 2, 3}).GetItems())
	})

	t.Run("can set inner Items", func(t *testing.T) {
		myList := list.New([]int{})
		myList.SetItems([]int{1, 2})
		assert.Equal(t, []int{1, 2}, myList.GetItems())
	})
}

func ExampleList_GetItems() {
	myList := list.New([]int{1, 2, 3})
	fmt.Println(myList.GetItems())
	// Output [1 2 3]
}

func ExampleList_SetItems() {
	myList := list.New([]int{})
	myList.SetItems([]int{1, 2})
	fmt.Println(myList.GetItems())
	// Output [1 2]
}

func BenchmarkList_GetItems(b *testing.B) {
	myList := list.New([]int{1, 2, 3})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myList.GetItems()
	}
}

func BenchmarkList_SetItems(b *testing.B) {
	myList := list.New([]int{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myList.SetItems([]int{1, 2, 3})
	}
}

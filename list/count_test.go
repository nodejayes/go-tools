package list_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/v2/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_Count(t *testing.T) {
	t.Run("0 when list is empty", func(t *testing.T) {
		myList := list.New(make([]int, 0))
		assert.Equal(t, 0, myList.Count())
	})
	t.Run("list [1 2 3] has 3 elements", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3})
		assert.Equal(t, 3, myList.Count())
	})
}

func ExampleList_Count() {
	myList := list.New([]int{1, 2, 3})
	fmt.Println(myList.Count())
	// Output: 3
}

func BenchmarkList_Count(b *testing.B) {
	myList := list.New([]int{1, 2, 3})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myList.Count()
	}
}

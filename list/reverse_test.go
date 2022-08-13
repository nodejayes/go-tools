package list_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_Reverse(t *testing.T) {
	t.Run("turn around [1 2 3]", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3})
		assert.Equal(t, []int{3, 2, 1}, myList.Reverse().GetItems())
	})
	t.Run("turn around two times results in original", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3})
		assert.Equal(t, []int{1, 2, 3}, myList.Reverse().Reverse().GetItems())
	})
}

func ExampleList_Reverse() {
	myList := list.New([]int{1, 2, 3})
	fmt.Println(myList.Reverse().GetItems())
	// Output: [3 2 1]
}

func BenchmarkList_Reverse(b *testing.B) {
	myList := list.New([]int{1, 2, 3})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myList.Reverse()
	}
}

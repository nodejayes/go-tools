package list_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFind(t *testing.T) {
	t.Run("", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3})
		assert.Equal(t, 2, myList.Find(func(item int, _ int, _ []int) bool {
			return item == 2
		}))
	})
}

func ExampleList_Find() {
	myList := list.New([]int{1, 2, 3})
	fmt.Println(myList.Find(func(item int, _ int, _ []int) bool {
		return item == 2
	}))
	// Output: 2
}

func BenchmarkList_Find(b *testing.B) {
	myList := list.New([]int{1, 2, 3})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myList.Find(func(item int, _ int, _ []int) bool {
			return item == 2
		})
	}
}

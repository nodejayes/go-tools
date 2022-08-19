package list_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/v3/compare"
	"github.com/nodejayes/go-tools/v3/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_Any(t *testing.T) {
	myList := list.New([]int{1, 2, 3})
	myEmptyList := list.New(make([]int, 0))
	t.Run("[1 2 3] has any", func(t *testing.T) {
		assert.True(t, myList.Any())
	})
	t.Run("[] has not any", func(t *testing.T) {
		assert.False(t, myEmptyList.Any())
	})
	t.Run("[1 2 3] has any 2", func(t *testing.T) {
		assert.True(t, myList.Any(func(item int, _ int, _ *list.List[int]) bool {
			return item == 2
		}))
	})
	t.Run("[1 2 3] has not any 5", func(t *testing.T) {
		assert.False(t, myList.Any(func(item int, _ int, _ *list.List[int]) bool {
			return item == 5
		}))
	})
}

func ExampleList_Any() {
	myList := list.New([]int{1, 2, 3})
	fmt.Println(myList.Any())
	fmt.Println(myList.Any(compare.FilterEquals(5)))
	// Output: true
	//false
}

func BenchmarkList_Any(b *testing.B) {
	myList := list.New([]int{1, 2, 3})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myList.Any()
	}
}

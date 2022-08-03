package list_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/compare"
	"github.com/nodejayes/go-tools/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFind(t *testing.T) {
	t.Run("find 2 in [1 2 3]", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3})
		assert.Equal(t, 2, myList.Find(compare.FilterEquals(2)))
	})

	t.Run("find nothing in empty list and return default", func(t *testing.T) {
		myList := list.New([]int{})
		assert.Equal(t, 0, myList.Find(compare.FilterEquals(2)))
	})
}

func ExampleList_Find() {
	myList := list.New([]int{1, 2, 3})
	fmt.Println(myList.Find(compare.FilterEquals(2)))
	// Output: 2
}

func BenchmarkList_Find(b *testing.B) {
	myList := list.New([]int{1, 2, 3})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myList.Find(compare.FilterEquals(2))
	}
}

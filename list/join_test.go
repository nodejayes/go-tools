package list_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_Join(t *testing.T) {
	t.Run("join [1 2 3 4 5] with #", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3, 4, 5})
		assert.Equal(t, "1#2#3#4#5", myList.Join("#"))
	})
	t.Run("join [1 2 3 4 5] with # custom formatter", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3, 4, 5})
		assert.Equal(t, "N: 1#N: 2#N: 3#N: 4#N: 5", myList.Join("#", func(item int) string {
			return fmt.Sprintf("N: %v", item)
		}))
	})
}

func ExampleList_Join() {
	myList := list.New([]int{1, 2, 3, 4, 5})
	fmt.Println(myList.Join("#"))
	// Output: 1#2#3#4#5
}

func BenchmarkList_Join(b *testing.B) {
	myList := list.New([]int{1, 2, 3, 4, 5})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myList.Join("#")
	}
}

package list_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_Copy(t *testing.T) {
	t.Run("copy [1 2 3]", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3})
		copyList := myList.Copy()
		myList.Clear()
		assert.Equal(t, 3, copyList.Count())
		assert.Equal(t, 0, myList.Count())
	})
}

func ExampleList_Copy() {
	myList := list.New([]int{1, 2, 3})
	copyList := myList.Copy()
	fmt.Println(copyList.GetItems())
	// Output: [1 2 3]
}

func BenchmarkList_Copy(b *testing.B) {
	myList := list.New([]int{1, 2, 3})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myList.Copy()
	}
}

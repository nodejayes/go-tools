package list_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/v2/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_ForSegment(t *testing.T) {
	t.Run("no segments when list has no item", func(t *testing.T) {
		myList := list.New([]int{})
		counter := 0
		myList.ForSegment(func(item1, item2 int) {
			counter++
		})
		assert.Equal(t, 0, counter)
	})
	t.Run("no segments when list has one item", func(t *testing.T) {
		myList := list.New([]int{1})
		counter := 0
		myList.ForSegment(func(item1, item2 int) {
			counter++
		})
		assert.Equal(t, 0, counter)
	})
	t.Run("one segment when list has two items", func(t *testing.T) {
		myList := list.New([]int{1, 2})
		counter := 0
		myList.ForSegment(func(item1, item2 int) {
			counter++
		})
		assert.Equal(t, 1, counter)
	})
}

func ExampleList_ForSegment() {
	myList := list.New([]int{1, 2})
	sum := 0
	myList.ForSegment(func(n1, n2 int) {
		sum += n1 + n2
	})
	fmt.Println(sum)
	// Output: 3
}

func BenchmarkList_ForSegment(b *testing.B) {
	myList := list.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	sum := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myList.ForSegment(func(n1, n2 int) {
			sum += n1 + n2
		})
	}
}

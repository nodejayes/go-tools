package list_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_Chunk(t *testing.T) {
	t.Run("[1 2 3 4 5] to [[1 2] [3 4] [5]]", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3, 4, 5})
		chunked := list.New(myList.Chunk(2))
		assert.Equal(t, 3, chunked.Count())

		part, err := chunked.ElementAt(0)
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 2}, part.GetItems())

		part, err = chunked.ElementAt(1)
		assert.NoError(t, err)
		assert.Equal(t, []int{3, 4}, part.GetItems())

		part, err = chunked.ElementAt(2)
		assert.NoError(t, err)
		assert.Equal(t, []int{5}, part.GetItems())
	})
	t.Run("[1 2 3 4 5] to [[1] [2] [3] [4] [5]]", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3, 4, 5})
		chunked := list.New(myList.Chunk(1))
		assert.Equal(t, 5, chunked.Count())

		part, err := chunked.ElementAt(0)
		assert.NoError(t, err)
		assert.Equal(t, []int{1}, part.GetItems())

		part, err = chunked.ElementAt(1)
		assert.NoError(t, err)
		assert.Equal(t, []int{2}, part.GetItems())

		part, err = chunked.ElementAt(2)
		assert.NoError(t, err)
		assert.Equal(t, []int{3}, part.GetItems())

		part, err = chunked.ElementAt(3)
		assert.NoError(t, err)
		assert.Equal(t, []int{4}, part.GetItems())

		part, err = chunked.ElementAt(4)
		assert.NoError(t, err)
		assert.Equal(t, []int{5}, part.GetItems())
	})
	t.Run("[1 2 3 4 5] to [1 2 3 4 5] when chunk size is 0", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3, 4, 5})
		chunked := list.New(myList.Chunk(0))
		assert.Equal(t, 1, chunked.Count())
		part, err := chunked.ElementAt(0)
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3, 4, 5}, part.GetItems())
	})
	t.Run("[1 2 3 4 5] to [1 2 3 4 5] when chunk size is -1", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3, 4, 5})
		chunked := list.New(myList.Chunk(-1))
		assert.Equal(t, 1, chunked.Count())
		part, err := chunked.ElementAt(0)
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3, 4, 5}, part.GetItems())
	})
}

func ExampleList_Chunk() {
	myList := list.New([]int{1, 2, 3, 4, 5})
	chunked := list.New(myList.Chunk(2))
	part1, _ := chunked.ElementAt(0)
	part2, _ := chunked.ElementAt(1)
	part3, _ := chunked.ElementAt(2)
	fmt.Println(part1.GetItems())
	fmt.Println(part2.GetItems())
	fmt.Println(part3.GetItems())
	// Output: [1 2]
	//[3 4]
	//[5]
}

func BenchmarkList_Chunk(b *testing.B) {
	myList := list.New([]int{1, 2, 3, 4, 5})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myList.Chunk(2)
	}
}

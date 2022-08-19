package list_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/v3/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_Pull(t *testing.T) {
	t.Run("pull 2 out of [1 2 3]", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3})
		val := myList.Pull(1)
		assert.Equal(t, 2, val)
		assert.Equal(t, []int{1, 3}, myList.GetItems())
	})
}

func ExampleList_Pull() {
	myList := list.New([]int{1, 2, 3})
	val := myList.Pull(1)
	fmt.Println(val)
	fmt.Println(myList.GetItems())
	// Output: 2
	//[1 3]
}

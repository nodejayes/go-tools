package list_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/v3/compare"
	"github.com/nodejayes/go-tools/v3/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindIndex(t *testing.T) {
	t.Run("find index of item in list", func(t *testing.T) {
		idx, err := list.New([]int{1, 2, 3}).FindIndex(compare.FilterEquals(1))
		assert.NoError(t, err)
		assert.Equal(t, 0, idx)
	})

	t.Run("returns error when not find", func(t *testing.T) {
		idx, err := list.New([]int{1, 2, 3}).FindIndex(compare.FilterEquals(5))
		assert.Error(t, err)
		assert.Equal(t, 0, idx)
	})
}

func ExampleList_FindIndex() {
	idx, err := list.New([]int{1, 2, 3}).FindIndex(compare.FilterEquals(1))
	fmt.Println(idx)

	idx, err = list.New([]int{1, 2, 3}).FindIndex(compare.FilterEquals(5))
	fmt.Println(err.Error())
	// Output: 0
	//item not found in list
}

func BenchmarkList_FindIndex(b *testing.B) {
	myList := list.New([]int{1, 2, 3})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = myList.FindIndex(compare.FilterEquals(1))
	}
}

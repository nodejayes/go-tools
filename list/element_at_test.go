package list_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/v3/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_ElementAt(t *testing.T) {
	t.Run("find a element at index 0 in list [1 2 3]", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3})
		element, err := myList.ElementAt(0)
		assert.NoError(t, err)
		assert.Equal(t, 1, element)
	})

	t.Run("index to high returns an error", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3})
		element, err := myList.ElementAt(3)
		assert.Error(t, err)
		assert.Equal(t, 0, element)
	})
}

func ExampleList_ElementAt() {
	myList := list.New([]int{1, 2, 3})
	// the error is nil
	element, err := myList.ElementAt(0)
	fmt.Println(fmt.Sprintf("%v %v", element, err))

	// the default value of the type was returned and a error
	element, err = myList.ElementAt(3)
	fmt.Println(fmt.Sprintf("%v %v", element, err.Error()))
	// Output: 1 <nil>
	//0 element with index 3 not in list. list has only 3 elements
}

func BenchmarkList_ElementAt(b *testing.B) {
	myList := list.New([]int{1, 2, 3})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = myList.ElementAt(0)
	}
}

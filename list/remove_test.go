package list_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/compare"
	"github.com/nodejayes/go-tools/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_Remove(t *testing.T) {
	t.Run("remove 5 from [1 2 3]", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3})
		myList.Remove(compare.FilterEquals(5))
		assert.Equal(t, []int{1, 2, 3}, myList.GetItems())
	})
	t.Run("remove 2 from [1 2 3]", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3})
		myList.Remove(compare.FilterEquals(2))
		assert.Equal(t, []int{1, 3}, myList.GetItems())
	})
	t.Run("remove custom strategy", func(t *testing.T) {
		myList := list.New([]testUser{{
			Name: "Paul",
			Age:  12,
		}, {
			Name: "Mina",
			Age:  18,
		}})
		myList.Remove(func(item testUser, _ int, _ list.List[testUser]) bool {
			return item.Age > 15
		})
		assert.Equal(t, []testUser{{
			Name: "Paul",
			Age:  12,
		}}, myList.GetItems())
	})
}

func ExampleList_Remove() {
	myList := list.New([]int{1, 2, 3})
	myList.Remove(compare.FilterEquals(2))
	fmt.Println(myList.GetItems())
	// Output: [1 3]
}

func BenchmarkList_Remove(b *testing.B) {
	myList := list.New([]int{1, 2, 3})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myList.Remove(compare.FilterEquals(2))
	}
}

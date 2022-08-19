package list_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/v3/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_Sort(t *testing.T) {
	t.Run("sort a List of integers", func(t *testing.T) {
		myList := list.New([]int{5, 1, 3, 2, 4})
		myList.Sort(func(n1, n2 int) bool {
			return n1 < n2
		})
		assert.Equal(t, []int{1, 2, 3, 4, 5}, myList.GetItems())
		myList.Sort(func(n1, n2 int) bool {
			return n1 > n2
		})
		assert.Equal(t, []int{5, 4, 3, 2, 1}, myList.GetItems())
	})
	t.Run("sort a List of Structs", func(t *testing.T) {
		myList := list.New([]testUser{{
			Name: "Paul",
			Age:  12,
		}, {
			Name: "Mina",
			Age:  10,
		}, {
			Name: "Axel",
			Age:  18,
		}, {
			Name: "Mina",
			Age:  18,
		}})
		myList.Sort(func(u1, u2 testUser) bool {
			return u1.Name < u2.Name || u1.Age > u2.Age
		})
		assert.Equal(t, []testUser{{
			Name: "Axel",
			Age:  18,
		}, {
			Name: "Mina",
			Age:  18,
		}, {
			Name: "Mina",
			Age:  10,
		}, {
			Name: "Paul",
			Age:  12,
		}}, myList.GetItems())
	})
}

func ExampleList_Sort() {
	myList := list.New([]int{5, 1, 3, 2, 4})
	myList.Sort(func(n1, n2 int) bool {
		return n1 < n2
	})
	fmt.Println(myList.GetItems())
	// Output: [1 2 3 4 5]
}

func BenchmarkList_Sort(b *testing.B) {
	myList := list.New([]int{5, 1, 3, 2, 4})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myList.Sort(func(n1, n2 int) bool {
			return n1 < n2
		})
	}
}

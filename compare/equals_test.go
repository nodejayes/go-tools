package compare_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/v3/compare"
	"github.com/nodejayes/go-tools/v3/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEquals(t *testing.T) {
	t.Run("bool", func(t *testing.T) {
		assert.True(t, compare.Equals(true, true))
		assert.False(t, compare.Equals(true, false))
	})

	t.Run("int", func(t *testing.T) {
		assert.True(t, compare.Equals(1, 1))
		assert.False(t, compare.Equals(1, 2))
	})

	t.Run("float", func(t *testing.T) {
		assert.True(t, compare.Equals(1.2, 1.2))
		assert.False(t, compare.Equals(1.2, 1.3))
	})

	t.Run("string", func(t *testing.T) {
		assert.True(t, compare.Equals("abc", "abc"))
		assert.False(t, compare.Equals("abc", "abcd"))
	})
}

func TestFilterEquals(t *testing.T) {
	t.Run("element at index 0 equals 1 in [1 2 3]", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3})
		el, _ := myList.ElementAt(0)
		assert.True(t, compare.FilterEquals(1)(el, 0, myList))
	})

	t.Run("element at index 1 not equals 1 in [1 2 3]", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3})
		el, _ := myList.ElementAt(1)
		assert.False(t, compare.FilterEquals(1)(el, 1, myList))
	})
}

func TestFilterNotEquals(t *testing.T) {
	t.Run("element at index 1 not equals 1 in [1 2 3]", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3})
		el, _ := myList.ElementAt(1)
		assert.True(t, compare.FilterNotEquals(1)(el, 1, myList))
	})

	t.Run("element at index 0 equals 1 in [1 2 3]", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3})
		el, _ := myList.ElementAt(0)
		assert.False(t, compare.FilterNotEquals(1)(el, 0, myList))
	})
}

func TestFilterTrue(t *testing.T) {
	myList := list.New([]int{1, 2, 3})
	t.Run("returns true", func(t *testing.T) {
		assert.True(t, compare.FilterTrue[int]()(1, 1, myList))
	})
}

func ExampleEquals() {
	fmt.Println(compare.Equals(1, 2))
	// Output false

	fmt.Println(compare.Equals(1, 1))
	// Output true
}

func BenchmarkEquals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		compare.Equals(1, 2)
	}
}

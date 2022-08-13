package list_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_Add(t *testing.T) {
	t.Run("add 1 to [3 2]", func(t *testing.T) {
		myList := list.New([]int{3, 2})
		myList.Add(1)
		assert.Equal(t, []int{3, 2, 1}, myList.GetItems())
	})
	t.Run("add 5 to [1 2 3]", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3})
		myList.AddIfNotExists(5)
		assert.Equal(t, []int{1, 2, 3, 5}, myList.GetItems())
	})
	t.Run("not add 1 to [1 2 3]", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3})
		myList.AddIfNotExists(1)
		assert.Equal(t, []int{1, 2, 3}, myList.GetItems())
	})
	t.Run("add Mina to [Paul Petra]", func(t *testing.T) {
		myList := list.New([]testUser{{
			Name: "Paul",
			Age:  12,
		}, {
			Name: "Petra",
			Age:  18,
		}})
		myList.AddIfNotExists(testUser{
			Name: "Mina",
			Age:  18,
		}, func(item testUser, _ int, _ list.List[testUser]) bool {
			return item.Name == "Mina"
		})
		assert.Equal(t, []testUser{{
			Name: "Paul",
			Age:  12,
		}, {
			Name: "Petra",
			Age:  18,
		}, {
			Name: "Mina",
			Age:  18,
		}}, myList.GetItems())
	})
	t.Run("not add Mina to [Paul Petra]", func(t *testing.T) {
		myList := list.New([]testUser{{
			Name: "Paul",
			Age:  12,
		}, {
			Name: "Petra",
			Age:  18,
		}})
		myList.AddIfNotExists(testUser{
			Name: "Mina",
			Age:  18,
		}, func(item testUser, _ int, _ list.List[testUser]) bool {
			return item.Age == 18
		})
		assert.Equal(t, []testUser{{
			Name: "Paul",
			Age:  12,
		}, {
			Name: "Petra",
			Age:  18,
		}}, myList.GetItems())
	})
	t.Run("add multiple items to List", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3})
		myList.AddRange(list.New([]int{1, 2, 3}))
		assert.Equal(t, []int{1, 2, 3, 1, 2, 3}, myList.GetItems())
	})
	t.Run("add if not exists multiple items to list", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3})
		myList.AddRangeIfNotExists(list.New([]int{1, 2, 3, 4}))
		assert.Equal(t, []int{1, 2, 3, 4}, myList.GetItems())
	})
}

func ExampleList_Add() {
	myList := list.New([]int{3, 2})
	myList.Add(1)
	fmt.Println(myList.GetItems())
	// Output: [3 2 1]
}

func ExampleList_AddIfNotExists() {
	myList := list.New([]int{1, 2, 3})
	myList.AddIfNotExists(2)
	fmt.Println(myList.GetItems())
	myList.AddIfNotExists(4)
	fmt.Println(myList.GetItems())
	// Output: [1 2 3]
	//[1 2 3 4]
}

func ExampleList_AddRange() {
	myList := list.New([]int{1, 2, 3})
	myList.AddRange(list.New([]int{1, 2, 3}))
	fmt.Println(myList.GetItems())
	// Output: [1 2 3 1 2 3]
}

func ExampleList_AddRangeIfNotExists() {
	myList := list.New([]int{1, 2, 3})
	myList.AddRangeIfNotExists(list.New([]int{1, 2, 3, 4}))
	fmt.Println(myList.GetItems())
	// Output: [1 2 3 4]
}

func BenchmarkList_Add(b *testing.B) {
	myList := list.New([]int{3, 2})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myList.Add(1)
	}
}

func BenchmarkList_AddIfNotExists(b *testing.B) {
	myList := list.New([]int{1, 2, 3})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myList.AddIfNotExists(2)
	}
}

func BenchmarkList_AddRange(b *testing.B) {
	myList := list.New([]int{1, 2, 3})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myList.AddRange(list.New([]int{1, 2, 3}))
	}
}

func BenchmarkList_AddRangeIfNotExists(b *testing.B) {
	myList := list.New([]int{1, 2, 3})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myList.AddRangeIfNotExists(list.New([]int{1, 2, 3, 4}))
	}
}

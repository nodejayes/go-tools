package list_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gitlab.com/morgenstern87/go-tools/list"
	"testing"
)

type (
	testUser struct {
		Name string
		Age  int
	}
)

func TestMap(t *testing.T) {
	t.Run("int to string", func(t *testing.T) {
		res := list.Map([]int{1, 2, 3}, func(item int, _ int, _ []int) string {
			return fmt.Sprintf("%v", item)
		})
		assert.Equal(t, []string{"1", "2", "3"}, res)
	})

	t.Run("float64 to string", func(t *testing.T) {
		res := list.Map([]float64{2.546, 1, 5.2}, func(item float64, _ int, _ []float64) string {
			return fmt.Sprintf("%0.2f", item)
		})
		assert.Equal(t, []string{"2.55", "1.00", "5.20"}, res)
	})

	t.Run("struct to string", func(t *testing.T) {
		res := list.Map([]testUser{
			{Name: "Tom", Age: 15},
			{Name: "Martina", Age: 26},
			{Name: "Paul", Age: 50},
			{Name: "Gerta", Age: 80},
		}, func(item testUser, _ int, _ []testUser) string {
			return fmt.Sprintf("%v (%v)", item.Name, item.Age)
		})
		assert.Equal(t, []string{"Tom (15)", "Martina (26)", "Paul (50)", "Gerta (80)"}, res)
	})

	t.Run("empty array returns empty array", func(t *testing.T) {
		res := list.Map([]int{}, func(item int, _ int, _ []int) string { return fmt.Sprintf("%v", item) })
		assert.Equal(t, []string{}, res)
	})
}

func ExampleMap() {
	fmt.Println(list.Map([]int{1, 2, 3}, func(item int, _ int, _ []int) string {
		return fmt.Sprintf("%v", item)
	}))
	// Output: [1 2 3]
}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list.Map([]int{1, 2, 3}, func(item int, _ int, _ []int) string {
			return fmt.Sprintf("%v", item)
		})
	}
}

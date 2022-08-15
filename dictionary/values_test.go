package dictionary_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/v2/dictionary"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDictionary_Values(t *testing.T) {
	t.Run("", func(t *testing.T) {
		myDic := dictionary.New(map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		})
		values := myDic.Values()
		values.Sort(func(v1, v2 int) bool {
			return v1 < v2
		})
		assert.Equal(t, []int{1, 2, 3}, values.GetItems())
	})
}

func ExampleDictionary_Values() {
	myDic := dictionary.New(map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	})
	values := myDic.Values()
	values.Sort(func(v1, v2 int) bool {
		return v1 < v2
	})
	fmt.Println(values.GetItems())
	// Output: [1 2 3]
}

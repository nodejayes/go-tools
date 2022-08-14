package dictionary_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/dictionary"
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
		assert.Equal(t, []int{1, 2, 3}, myDic.Values().GetItems())
	})
}

func ExampleDictionary_Values() {
	myDic := dictionary.New(map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	})
	fmt.Println(myDic.Values().GetItems())
	// Output: [1 2 3]
}

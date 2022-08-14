package dictionary_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/dictionary"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDictionary_Keys(t *testing.T) {
	t.Run("get Dictionary Keys [a b c]", func(t *testing.T) {
		myDic := dictionary.New(map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		})
		assert.Equal(t, []string{"a", "b", "c"}, myDic.Keys().GetItems())
	})
}

func ExampleDictionary_Keys() {
	myDic := dictionary.New(map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	})
	fmt.Println(myDic.Keys().GetItems())
	// Output: [a b c]
}

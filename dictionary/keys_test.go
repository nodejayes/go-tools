package dictionary_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/v2/dictionary"
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
		keys := myDic.Keys()
		keys.Sort(func(k1, k2 string) bool {
			return k1 < k2
		})
		assert.Equal(t, []string{"a", "b", "c"}, keys.GetItems())
	})
}

func ExampleDictionary_Keys() {
	myDic := dictionary.New(map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	})
	keys := myDic.Keys()
	keys.Sort(func(k1, k2 string) bool {
		return k1 < k2
	})
	fmt.Println(keys.GetItems())
	// Output: [a b c]
}

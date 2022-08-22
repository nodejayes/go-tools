package dictionary_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/v3/dictionary"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDictionary_HasValue(t *testing.T) {
	t.Run("find a value in the dictionary", func(t *testing.T) {
		myDic := dictionary.New(map[string]int{
			"a": 1,
		})
		assert.True(t, myDic.HasValue(1))
	})
	t.Run("not find a value in the dictionary", func(t *testing.T) {
		myDic := dictionary.New(map[string]int{
			"a": 1,
		})
		assert.False(t, myDic.HasValue(2))
	})
}

func ExampleDictionary_HasValue() {
	myDic := dictionary.New(map[string]int{
		"a": 1,
	})
	fmt.Println(myDic.HasValue(1))
	// Output: true
}

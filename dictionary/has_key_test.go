package dictionary_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/v3/dictionary"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDictionary_HasKey(t *testing.T) {
	t.Run("find a existing key", func(t *testing.T) {
		myDic := dictionary.New(map[string]int{
			"a": 1,
		})
		assert.True(t, myDic.HasKey("a"))
	})
	t.Run("not find a not existing key", func(t *testing.T) {
		myDic := dictionary.New(map[string]int{
			"a": 1,
		})
		assert.False(t, myDic.HasKey("b"))
	})
}

func ExampleDictionary_HasKey() {
	myDic := dictionary.New(map[string]int{
		"a": 1,
	})
	fmt.Println(myDic.HasKey("a"))
	// Output: true
}

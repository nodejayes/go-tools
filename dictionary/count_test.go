package dictionary_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/v3/dictionary"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDictionary_Count(t *testing.T) {
	t.Run("empty Dictionary", func(t *testing.T) {
		myDic := dictionary.New(make(map[string]int))
		assert.Equal(t, 0, myDic.Count())
	})
	t.Run("filled Dictionary", func(t *testing.T) {
		myDic := dictionary.New(map[string]int{"a": 1})
		assert.Equal(t, 1, myDic.Count())
	})
}

func ExampleDictionary_Count() {
	myDic := dictionary.New(map[string]int{"a": 1})
	fmt.Println(myDic.Count())
	// Output: 1
}

package dictionary_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/v3/dictionary"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDictionary_GetMap(t *testing.T) {
	t.Run("get the internal map", func(t *testing.T) {
		myDic := dictionary.New(map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		})
		assert.Equal(t, map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		}, myDic.GetMap())
	})
}

func ExampleDictionary_GetMap() {
	myDic := dictionary.New(map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	})
	fmt.Println(myDic.GetMap())
	// Output: map[a:1 b:2 c:3]
}

package dictionary_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/dictionary"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDictionary_SetMap(t *testing.T) {
	t.Run("set the internal map of the Dictionary", func(t *testing.T) {
		myDic := dictionary.New(map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		})
		myDic.SetMap(map[string]int{
			"d": 4,
		})
		assert.Equal(t, map[string]int{
			"d": 4,
		}, myDic.GetMap())
	})
}

func ExampleDictionary_SetMap() {
	myDic := dictionary.New(map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	})
	myDic.SetMap(map[string]int{
		"d": 4,
	})
	fmt.Println(myDic.GetMap())
	// Output: map[d:4]
}

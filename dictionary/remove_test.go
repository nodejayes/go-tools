package dictionary_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/v3/dictionary"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDictionary_Remove(t *testing.T) {
	t.Run("remove a value when key exists", func(t *testing.T) {
		myDic := dictionary.New(map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		})
		assert.Equal(t, 3, myDic.Keys().Count())
		myDic.Remove("a")
		assert.Equal(t, 2, myDic.Keys().Count())
		_, f := myDic.GetValue("a")
		assert.Error(t, f)
	})
	t.Run("not remove when key not exists", func(t *testing.T) {
		myDic := dictionary.New(map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		})
		assert.Equal(t, 3, myDic.Keys().Count())
		myDic.Remove("x")
		assert.Equal(t, 3, myDic.Keys().Count())
		_, f := myDic.GetValue("x")
		assert.Error(t, f)
	})
}

func ExampleDictionary_Remove() {
	myDic := dictionary.New(map[string]int{
		"a": 1,
		"b": 2,
	})
	myDic.Remove("a")
	fmt.Println(myDic.String())
	// Output: {b:2}
}

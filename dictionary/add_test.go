package dictionary_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/v3/dictionary"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDictionary_Add(t *testing.T) {
	t.Run("add a not existing key", func(t *testing.T) {
		myDic := dictionary.New(map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		})
		myDic.Add("d", 4)
		assert.Equal(t, 4, myDic.Keys().Count())
		v, f := myDic.GetValue("d")
		assert.Nil(t, f)
		assert.Equal(t, 4, v)
	})
	t.Run("overwrite value that exists", func(t *testing.T) {
		myDic := dictionary.New(map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		})
		myDic.Add("a", 5)
		assert.Equal(t, 3, myDic.Keys().Count())
		v, f := myDic.GetValue("a")
		assert.Nil(t, f)
		assert.Equal(t, 5, v)
	})
	t.Run("add nil map", func(t *testing.T) {
		myDic := dictionary.New[string, int](nil)
		myDic.Add("a", 5)
		assert.Equal(t, 1, myDic.Keys().Count())
		v, f := myDic.GetValue("a")
		assert.Nil(t, f)
		assert.Equal(t, 5, v)
	})
}

func ExampleDictionary_Add() {
	myDic := dictionary.New(map[string]int{})
	myDic.Add("d", 4)
	fmt.Println(myDic.String())
	// Output: {d:4}
}

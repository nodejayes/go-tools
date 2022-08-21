package dictionary_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/v3/dictionary"
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
		values := myDic.Values()
		values.Sort(func(v1, v2 int) bool {
			return v1 < v2
		})
		assert.Equal(t, []int{1, 2, 3}, values.GetItems())
	})
	t.Run("get the value of a key that exists", func(t *testing.T) {
		myDic := dictionary.New(map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		})
		v, f := myDic.GetValue("a")
		assert.Nil(t, f)
		assert.Equal(t, 1, v)
	})
	t.Run("get error when key not exists", func(t *testing.T) {
		myDic := dictionary.New(map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		})
		v, f := myDic.GetValue("x")
		assert.Error(t, f)
		assert.Equal(t, 0, v)
	})
}

func ExampleDictionary_Values() {
	myDic := dictionary.New(map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	})
	values := myDic.Values()
	values.Sort(func(v1, v2 int) bool {
		return v1 < v2
	})
	fmt.Println(values.GetItems())
	// Output: [1 2 3]
}

func ExampleDictionary_GetValue() {
	myDic := dictionary.New(map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	})
	v, f := myDic.GetValue("a")
	fmt.Println(fmt.Sprintf("%v %v", v, f))
	// Output: 1 <nil>
}

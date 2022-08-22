package dictionary_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/v3/dictionary"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDictionary_Clone(t *testing.T) {
	t.Run("clone a dictionary", func(t *testing.T) {
		myDic := dictionary.New(map[string]int{
			"a": 1,
		})
		mySecondDic := myDic.Clone()
		mySecondDic.Add("b", 2)
		assert.True(t, myDic.HasKey("a"))
		assert.False(t, myDic.HasKey("b"))
		assert.True(t, mySecondDic.HasKey("a"))
		assert.True(t, mySecondDic.HasKey("b"))
		v, _ := mySecondDic.GetValue("b")
		assert.Equal(t, 2, v)
	})
}

func ExampleDictionary_Clone() {
	myDic := dictionary.New(map[string]int{
		"a": 1,
	})
	mySecondDic := myDic.Clone()
	mySecondDic.Add("b", 2)
	fmt.Println(mySecondDic.HasKey("b"))
	fmt.Println(myDic.HasKey("b"))
	// Output: true
	//false
}

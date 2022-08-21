package dictionary_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/v3/dictionary"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("create a new Dictionary", func(t *testing.T) {
		myDic := dictionary.New(map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		})
		assert.NotNil(t, myDic)
	})
	t.Run("implements the Stringer Interface", func(t *testing.T) {
		myDic := dictionary.New(map[string]int{
			"a": 1,
		})
		assert.Equal(t, "{a:1}", myDic.String())
	})
}

func ExampleNew() {
	myDic := dictionary.New(map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	})
	fmt.Println(myDic)
	// Output: {map[a:1 b:2 c:3]}
}

func ExampleDictionary_String() {
	myDic := dictionary.New(map[string]int{
		"a": 1,
	})
	fmt.Println(myDic.String())
	// Output: {a:1}
}

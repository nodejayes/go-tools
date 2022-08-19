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
}

func ExampleNew() {
	myDic := dictionary.New[string, int](map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	})
	fmt.Println(myDic)
	// Output: {map[a:1 b:2 c:3]}
}

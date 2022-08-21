package dictionary_test

import (
	"github.com/nodejayes/go-tools/v3/dictionary"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDictionary_Clear(t *testing.T) {
	t.Run("clear a existing dictionary", func(t *testing.T) {
		myDic := dictionary.New(map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		})
		assert.Equal(t, 3, myDic.Keys().Count())
		myDic.Clear()
		assert.Equal(t, 0, myDic.Keys().Count())
	})
}

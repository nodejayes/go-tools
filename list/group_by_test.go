package list_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/v3/dictionary"
	"github.com/nodejayes/go-tools/v3/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_GroupBy(t *testing.T) {
	t.Run("group the list", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13})
		res := dictionary.New(myList.GroupBy(func(v int) string {
			return fmt.Sprintf("%v", v%2)
		}))
		assert.True(t, res.HasKey("0"))
	})
}

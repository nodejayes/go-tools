package counter_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/v3/counter"
	"github.com/nodejayes/go-tools/v3/fault"
	"github.com/nodejayes/go-tools/v3/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("create a new counter", func(t *testing.T) {
		c := counter.New(0)
		assert.Equal(t, 0, c.GetValue())
	})
	t.Run("can use in for loop", func(t *testing.T) {
		myList := list.New([]int{1, 2, 3, 4})
		c := 0
		for i := counter.New(0); i.GetValue() < myList.Count(); i.Increment() {
			var (
				v1 int
				v2 int
				f  *fault.Fault
			)
			v1, f = myList.ElementAt(i.GetValue())
			assert.Nil(t, f)
			v2, f = myList.ElementAt(c)
			assert.Nil(t, f)
			assert.Equal(t, v2, v1)
			c++
		}
	})
}

func ExampleNew() {
	c := counter.New(0)
	fmt.Println(c.GetValue())
	// Output: 0
}

package counter_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/v3/counter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCounter_Decrement(t *testing.T) {
	t.Run("decrement counter by 1", func(t *testing.T) {
		c := counter.New(2)
		c.Decrement(1)
		assert.Equal(t, 1, c.GetValue())
	})
	t.Run("decrement counter by default", func(t *testing.T) {
		c := counter.New(2)
		c.Decrement()
		assert.Equal(t, 1, c.GetValue())
	})
	t.Run("decrement counter by 2", func(t *testing.T) {
		c := counter.New(4)
		c.Decrement(2)
		assert.Equal(t, 2, c.GetValue())
	})
	t.Run("0 is the min value of the counter by 4", func(t *testing.T) {
		c := counter.New()
		c.Decrement(4)
		assert.Equal(t, 0, c.GetValue())
	})
	t.Run("0 is the min value of the counter by default", func(t *testing.T) {
		c := counter.New()
		c.Decrement()
		assert.Equal(t, 0, c.GetValue())
	})
}

func ExampleCounter_Decrement() {
	c := counter.New(5)
	c.Decrement()
	fmt.Println(c.GetValue())
	// Output: 4
}

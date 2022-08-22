package counter_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/v3/counter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCounter_Increment(t *testing.T) {
	t.Run("increment by 1", func(t *testing.T) {
		c := counter.New()
		c.Increment(1)
		assert.Equal(t, 1, c.GetValue())
	})
	t.Run("increment by default", func(t *testing.T) {
		c := counter.New()
		c.Increment()
		assert.Equal(t, 1, c.GetValue())
	})
	t.Run("increment by 2", func(t *testing.T) {
		c := counter.New()
		c.Increment(2)
		assert.Equal(t, 2, c.GetValue())
	})
}

func ExampleCounter_Increment() {
	c := counter.New()
	c.Increment()
	fmt.Println(c.GetValue())
	// Output: 1
}

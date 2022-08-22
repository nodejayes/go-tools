package counter_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/v3/counter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCounter_Reset(t *testing.T) {
	t.Run("reset the counter", func(t *testing.T) {
		c := counter.New(5)
		assert.Equal(t, 5, c.GetValue())
		c.Reset()
		assert.Equal(t, 0, c.GetValue())
	})
}

func ExampleCounter_Reset() {
	c := counter.New(5)
	c.Reset()
	fmt.Println(c.GetValue())
	// Output: 0
}

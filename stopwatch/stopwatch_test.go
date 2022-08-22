package stopwatch_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/v3/stopwatch"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("create new stopwatch", func(t *testing.T) {
		sw := stopwatch.New()
		assert.NotNil(t, sw)
	})
}

func ExampleNew() {
	sw := stopwatch.New()
	fmt.Println(sw)
	// Output: &{0 map[]}
}

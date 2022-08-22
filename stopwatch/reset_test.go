package stopwatch_test

import (
	"github.com/nodejayes/go-tools/v3/stopwatch"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStopwatch_Reset(t *testing.T) {
	t.Run("can reset stopwatch", func(t *testing.T) {
		sw := stopwatch.New()
		sw.Reset()
		assert.Equal(t, 0.0, sw.ElapsedMs())
	})
}

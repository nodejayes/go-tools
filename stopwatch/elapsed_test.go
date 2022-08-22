package stopwatch_test

import (
	"github.com/nodejayes/go-tools/v3/stopwatch"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStopwatch_ElapsedMs(t *testing.T) {
	t.Run("get elapsed ms", func(t *testing.T) {
		sw := stopwatch.New()
		assert.True(t, sw.ElapsedMs() > 0.0)
	})
	t.Run("get elapsed ms string", func(t *testing.T) {
		sw := stopwatch.New()
		assert.True(t, len(sw.ElapsedMsAsString()) > 0)
	})
}

func TestStopwatch_PrintElapsedMs(t *testing.T) {
	sw := stopwatch.New()
	sw.PrintElapsedMs("Test1", true)
	sw.PrintElapsedMs("Test2", false)
	// Output:
}

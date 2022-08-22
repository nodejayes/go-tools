package stopwatch_test

import (
	"github.com/nodejayes/go-tools/v3/stopwatch"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestStopwatch_StartSection(t *testing.T) {
	t.Run("start a section", func(t *testing.T) {
		sw := stopwatch.New()
		sw.StartSection("Test")
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Millisecond)
		}
		assert.True(t, sw.StopSection("Test") > 0.0)
	})
}

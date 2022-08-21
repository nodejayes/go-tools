package fault_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/v3/fault"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("create a new error", func(t *testing.T) {
		err := fault.NewError("Test")
		assert.True(t, err.Is(fault.Error))
		assert.True(t, err.OneOf(fault.Error, fault.Warning, fault.Info))
		assert.Equal(t, "Test", err.Error())
	})
	t.Run("create a new warning", func(t *testing.T) {
		err := fault.NewWarning("Test")
		assert.True(t, err.Is(fault.Warning))
		assert.True(t, err.OneOf(fault.Error, fault.Warning, fault.Info))
		assert.Equal(t, "Test", err.Error())
	})
	t.Run("create a new info", func(t *testing.T) {
		err := fault.NewInfo("Test")
		assert.True(t, err.Is(fault.Info))
		assert.True(t, err.OneOf(fault.Error, fault.Warning, fault.Info))
		assert.Equal(t, "Test", err.Error())
	})
	t.Run("create a new debug", func(t *testing.T) {
		err := fault.NewDebug("Test")
		assert.True(t, err.Is(fault.Debug))
		assert.False(t, err.OneOf(fault.Error, fault.Warning, fault.Info))
		assert.Equal(t, "Test", err.Error())
	})
	t.Run("implement Stringer interface", func(t *testing.T) {
		err := fault.NewDebug("Test")
		assert.Equal(t, "Test", err.String())
	})
	t.Run("false by unknown type", func(t *testing.T) {
		err := fault.NewDebug("Test")
		assert.False(t, err.Is(9999))
	})
}

func ExampleNewError() {
	err := fault.NewError("Test")
	fmt.Println(err.Error())
	// Output: Test
}

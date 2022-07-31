package compare_test

import (
	"fmt"
	"github.com/nodejayes/go-tools/compare"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEquals(t *testing.T) {
	t.Run("bool", func(t *testing.T) {
		assert.True(t, compare.Equals(true, true))
		assert.False(t, compare.Equals(true, false))
	})

	t.Run("int", func(t *testing.T) {
		assert.True(t, compare.Equals(1, 1))
		assert.False(t, compare.Equals(1, 2))
	})

	t.Run("float", func(t *testing.T) {
		assert.True(t, compare.Equals(1.2, 1.2))
		assert.False(t, compare.Equals(1.2, 1.3))
	})

	t.Run("string", func(t *testing.T) {
		assert.True(t, compare.Equals("abc", "abc"))
		assert.False(t, compare.Equals("abc", "abcd"))
	})
}

func ExampleEquals() {
	fmt.Println(compare.Equals(1, 2))
	// Output false

	fmt.Println(compare.Equals(1, 1))
	// Output true
}

func BenchmarkEquals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		compare.Equals(1, 2)
	}
}

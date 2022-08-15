package async_test

import (
	"github.com/nodejayes/go-tools/v2/async"
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
)

func runAdd(t *testing.T, cpus, items int) {
	res := make([]int, items)
	in := make(chan []int, items)
	err := async.StartWorkers(cpus, in, func(in chan []int) {
		for i := 0; i < items; i++ {
			v := []int{i, 1}
			in <- v
		}
	}, func(params []int) {
		res[params[0]] = params[1] + params[1]
	})
	assert.NoError(t, err)
	for _, v := range res {
		assert.Equal(t, 2, v)
	}
}

func TestStartWorkers(t *testing.T) {
	t.Run("run with 3 CPUs", func(t *testing.T) {
		runAdd(t, 3, 100)
	})

	t.Run("run with os CPUs", func(t *testing.T) {
		runAdd(t, runtime.NumCPU(), 100)
	})

	t.Run("run with 100 CPUs limit to os max cpus", func(t *testing.T) {
		runAdd(t, 100, 100)
	})
}

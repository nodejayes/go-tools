package async

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"
)

// StartWorkers start async workers for the number of parameter n parallel
func StartWorkers[T any](n int, in chan T, fill func(in chan T), unitOfWork func(param T)) error {
	if n > runtime.NumCPU() {
		n = runtime.NumCPU()
	}
	var err error
	errCounter := 0
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				if rec := recover(); rec != nil {
					errCounter++
					err = fmt.Errorf("%v => %v", rec, string(debug.Stack()))
					wg.Done()
					return
				}
				wg.Done()
			}()
			for p := range in {
				if errCounter > 0 {
					continue
				}
				unitOfWork(p)
			}
		}()
	}

	fill(in)
	close(in)
	wg.Wait()
	return err
}

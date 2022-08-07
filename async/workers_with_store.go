package async

import (
	"fmt"
	"runtime/debug"
	"sync"
)

func StartWorkersWithStore[T any, K any](n int, in chan T, fill func(in chan T), unitOfWork func(param T, store K)) error {
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
			var store K
			for p := range in {
				if errCounter > 0 {
					continue
				}
				unitOfWork(p, store)
			}
		}()
	}

	fill(in)
	close(in)
	wg.Wait()
	return err
}

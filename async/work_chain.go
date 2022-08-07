package async

import (
	"fmt"
	"runtime/debug"
	"sync"
)

func WorkChain(n int, unitsOfWork []func() error) error {
	var err error
	errCounter := 0
	var wg sync.WaitGroup
	in := make(chan func() error, len(unitsOfWork))

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
				tmpErr := p()
				if tmpErr != nil {
					errCounter++
					err = tmpErr
				}
			}
		}()
	}

	for _, unitOfWork := range unitsOfWork {
		in <- unitOfWork
	}
	close(in)
	wg.Wait()

	return err
}

package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) error {
	if len(tasks) == 0 || n == 0 {
		return nil
	}
	if n < 0 || m <= 0 {
		return ErrErrorsLimitExceeded
	}
	chWorkers := make(chan Task)

	wg := sync.WaitGroup{}
	wg.Add(n)

	var err error
	var errorsCount int32
	atomic.StoreInt32(&errorsCount, 0)

	for i := 0; i < n; i++ {
		go RunWorker(chWorkers, &wg, &errorsCount)
	}

	for _, v := range tasks {
		chWorkers <- v
		if atomic.LoadInt32(&errorsCount) >= int32(m) {
			err = ErrErrorsLimitExceeded
			break
		}
	}
	close(chWorkers)

	wg.Wait()
	return err
}

func RunWorker(task <-chan Task, wg *sync.WaitGroup, errorsCount *int32) {
	defer wg.Done()
	for v := range task {
		if v() != nil {
			atomic.AddInt32(errorsCount, 1)
		}
	}
}

package hw05parallelexecution

import (
	"errors"
	"sync"
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
	defer wg.Wait()

	var errorsCount int
	var err error

	for i := 0; i < n; i++ {
		go RunWorker(chWorkers, &wg, &errorsCount)
	}

	for _, v := range tasks {
		chWorkers <- v
		if errorsCount >= m {
			err = ErrErrorsLimitExceeded
			break
		}
	}
	close(chWorkers)

	return err
}

func RunWorker(task <-chan Task, wg *sync.WaitGroup, errorsCount *int) {
	defer wg.Done()
	for v := range task {
		if v() != nil {
			*errorsCount++
		}
	}
}

package worker

import (
	"sync"
)

var waitgroup sync.WaitGroup

type Worker struct {
	MaxWorker int
}

//

func NewWorker(maxWorker int) Worker {
	return Worker{
		MaxWorker: maxWorker,
	}
}

func (w Worker) Start() {
	waitgroup.Add(w.MaxWorker)
}

func (w Worker) Stop() {
	defer waitgroup.Done()
}

func (w Worker) Wait() {
	waitgroup.Wait()
}

func (w Worker) Job() {
	w.Start()
}

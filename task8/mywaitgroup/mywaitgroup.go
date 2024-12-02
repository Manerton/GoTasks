package mywaitgroup

type MyWaitGroup struct {
	semaphore chan struct{}
	done      chan struct{}
}

func InitMyWaitGroup() *MyWaitGroup {
	return &MyWaitGroup{
		semaphore: make(chan struct{}, 1),
		done:      make(chan struct{}),
	}
}

func (wg *MyWaitGroup) AddOne() {
	wg.semaphore <- struct{}{}
}

func (wg *MyWaitGroup) Done() {
	select {
	case <-wg.semaphore:
	default:
		panic("done more than add")
	}
	if len(wg.semaphore) == 0 {
		close(wg.done)
	}
}

func (wg *MyWaitGroup) Wait() {
	<-wg.done
}

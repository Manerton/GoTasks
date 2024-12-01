package functions

import "sync"

func MergeChannel(channels ...<-chan int) <-chan int {
	mergeChannel := make(chan int)

	wg := sync.WaitGroup{}

	mergeGoroutine := func(ch <-chan int) {
		defer wg.Done()
		for num := range ch {
			mergeChannel <- num
		}
	}

	for _, channel := range channels {
		wg.Add(1)
		go mergeGoroutine(channel)
	}

	go func() {
		// wait all for close mergeChannel
		wg.Wait()
		close(mergeChannel)
	}()

	return mergeChannel
}

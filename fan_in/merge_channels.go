package fan_in

import "sync"

// MergeChannels - принимает несколько каналов на вход и объединяет их в один
// Fan-in и merge channels синонимы
func MergeChannels(channels ...<-chan int) <-chan int {
	wg := &sync.WaitGroup{}
	wg.Add(len(channels))
	fanIn := make(chan int)

	for _, ch := range channels {
		go func(ch <-chan int, wg *sync.WaitGroup) {
			for {
				val, ok := <-ch
				if !ok {
					wg.Done()
					break
				}
				fanIn <- val
			}
		}(ch, wg)
	}
	go func(wg *sync.WaitGroup) {
		wg.Wait()
		close(fanIn)
	}(wg)
	return fanIn
}

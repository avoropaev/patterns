package singleton

import "sync"

func Example() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			getInstance()
		}()
	}

	wg.Wait()
}

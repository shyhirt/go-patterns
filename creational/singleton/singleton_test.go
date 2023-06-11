package singleton

import (
	"sync"
	"testing"
)

func Test_getSingleton(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			getSingleton()
			wg.Done()
		}()
	}
	wg.Wait()
}

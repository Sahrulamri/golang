package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
	"sync/atomic"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		go func() {
			group.Add(1)
			defer group.Done()
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}

	time.Sleep(5 * time.Second)
	group.Wait()
	fmt.Println("Counter", x)
}
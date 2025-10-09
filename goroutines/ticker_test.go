package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TickerTest(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	group := sync.WaitGroup{}
	group.Add(1)

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
			group.Done()
			ticker.Stop()
		}
	}()
}



func TestTickerStop(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	group := sync.WaitGroup{}
	group.Add(1)

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
			group.Done()
			ticker.Stop()
		}
	}()

	group.Wait()
	fmt.Println("Ticker stopped")
	}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for time := range channel {
		fmt.Println("Tick at", time)
		break
	}
}
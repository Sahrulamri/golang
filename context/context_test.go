package context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
	
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println("Background Context:", background)

	todo := context.TODO()
	fmt.Println("TODO Context:", todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "username", "john_doe")
	contextC := context.WithValue(contextA, "request_id", "12345")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println("contextA username:", contextA.Value("username"))
	fmt.Println("contextA request_id:", contextA.Value("request_id"))
	fmt.Println("contextB username:", contextB.Value("username"))
	fmt.Println("contextC request_id:", contextC.Value("request_id"))
	fmt.Println("contextD d:", contextD.Value("d"))
	fmt.Println("contextD username:", contextD.Value("username"))
	fmt.Println("contextE e:", contextE.Value("e"))
	fmt.Println("contextE username:", contextE.Value("username"))
	fmt.Println("contextF f:", contextF.Value("f"))
	fmt.Println("contextF request_id:", contextF.Value("request_id"))
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutines before:", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter:", n)
		if n == 10 {
			break
		}
	}
	cancel()

	time.Sleep(2 * time.Second)

	fmt.Println("Cancel Context")
	fmt.Println("Total Goroutines after:", runtime.NumGoroutine())
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total Goroutines before:", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter:", n)
	}
	// cancel()

	time.Sleep(2 * time.Second)

	fmt.Println("Cancel Context")
	fmt.Println("Total Goroutines after:", runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutines before:", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter:", n)
	}
	// cancel()

	time.Sleep(2 * time.Second)

	fmt.Println("Cancel Context")
	fmt.Println("Total Goroutines after:", runtime.NumGoroutine())
}
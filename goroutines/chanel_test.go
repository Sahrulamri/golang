package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	c := make(chan int)
	c <- 42
	fmt.Println(<-c)

	var data int = <-c
	fmt.Println(data)

	defer close(c)

	go func(){
		time.Sleep(1 * time.Second)
		c <- 42
		fmt.Println("Selesai")
	} ()

	time.Sleep(2 * time.Second)
	
}
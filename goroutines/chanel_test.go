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

	go func() {
		time.Sleep(1 * time.Second)
		c <- 42
		fmt.Println("Selesai")
	}()

	time.Sleep(2 * time.Second)

}

func GiveMeResponse(Channel chan string) {
	time.Sleep(2 * time.Second)
	Channel <- "Haii Golang"
}

func TestChannelAsParameter(t *testing.T) {
	c := make(chan string)
	defer close(c)
	
	go GiveMeResponse(c)

	data := <-c
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Haii Golang"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	c := make(chan string)
	

	go OnlyIn(c)
	go OnlyOut(c)

	time.Sleep(3 * time.Second)
	close(c)
}

func TestBufferedChannel(t *testing.T) {
	c := make(chan string, 3)
	defer close(c)

	c <- "Amri"
	fmt.Println("1. Selesai Mengirim Data ke Channel")
	c <- "Dwi"
	fmt.Println("2. Selesai Mengirim Data ke Channel")
	c <- "Saputra"
	fmt.Println("3. Selesai Mengirim Data ke Channel")

	fmt.Println(<-c)
	fmt.Println("1. Selesai Menerima Data dari Channel")
	fmt.Println(<-c)
	fmt.Println("2. Selesai Menerima Data dari Channel")
	fmt.Println(<-c)
	fmt.Println("3. Selesai Menerima Data dari Channel")
	}

	func TestRangeChannel(t *testing.T) {
		c := make(chan string)
		
		go func() {
			for i := 0; i < 10; i++ {
				c <- fmt.Sprintf("Data ke %d", i)
			}
			close(c)
		}()

		for data := range c {
			fmt.Println("Menerima data ",data)
		}
		fmt.Println("Selesai")
	}

func TestSelectChannel(t *testing.T) {
	c1 := make(chan string)
	c2 := make(chan string)
	defer close(c1)
	defer close(c2)

	go GiveMeResponse(c1)
	go GiveMeResponse(c2)

	counter := 0
LOOP:
	for {
		select {
		case data := <-c1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-c2:
			fmt.Println("Data dari channel 2", data)
			counter++
		}

		if counter == 2 {
			fmt.Println("Selesai")
			break LOOP
		}
	}
}

func TestDefaultSelectChannel(t *testing.T) {
	c1 := make(chan string)
	c2 := make(chan string)
	defer close(c1)
	defer close(c2)

	go GiveMeResponse(c1)
	go GiveMeResponse(c2)

	counter := 0
LOOP:
	for {
		select {
		case data := <-c1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-c2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}

		if counter == 2 {
			fmt.Println("Selesai")
			break LOOP
		}
	}
}


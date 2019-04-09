package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func service1(c chan string) {
	time.Sleep(3 * time.Second)
	c <- "Hello from service1"
}

func service2(c chan string) {
	time.Sleep(5 * time.Second)
	c <- "Hello from service 2"
}

func main() {
	fmt.Println("main() started", time.Since(start))
	chan1 := make(chan string)
	chan2 := make(chan string)

	go service(chan1)
	go service(chan2)

	select {
	case res := <-chan1:
		fmt.Println("Response form service 1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from service 2", res, time.Since(start))
	case <-time.After(2 * time.Second):
		fmt.Println("No responce received", time.Since(start))

	}
	fmt.Println("main() stopped", time.Since(start))
}

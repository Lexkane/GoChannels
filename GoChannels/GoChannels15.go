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
	fmt.Println("service1() started", time.Since(start))
	c <- "Hello from service1"
}

func service2(c chan string) {
	fmt.Println("service 2() started", time.Since(start))
	c <- "Hello from service2"
}

func main() {
	fmt.Println("main() started", time.Since(start))

	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	time.Sleep(3 * time.Second)

	select {
	case res := <-chan1:
		fmt.Println("Responce from service1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Responce from service 2", res, time.Since(start))
	default:
		fmt.Println("No responce received", time.Since(start))
	}
	fmt.Println("main() stopped", time.Since(start))
}

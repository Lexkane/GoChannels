package main

import "fmt"

func service(c chan string) {
	c <- "response"
}

//error all goroutines asleep

func main() {
	fmt.Println("main() started")
	var chan1 chan string
	go service(chan1)
	select {
	case res := <-chan1:
		fmt.Println("Response from chan1", res)
	}
	fmt.Println("main() stopped")

}

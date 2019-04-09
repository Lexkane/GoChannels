package main

import "fmt"

func service(c chan string) {
	c <- "responce"
}

func main() {
	fmt.Println("main() started ")

	var chan1 chan string
	go service(chan1)

	select {
	case res := <-chan1:
		fmt.Println("Response from chan1", res)
	default:
		fmt.Println("No response")
	}
	fmt.Println("main() stopped")
}

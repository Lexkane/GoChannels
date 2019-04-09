package main

import "fmt"

func service() {
	fmt.Println("Hello from service")
}

//deadlock example
func main() {
	fmt.Println("main() started")
	go service()

	select {}
	fmt.Println("main() stopped")
}

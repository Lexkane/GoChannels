package main

import "fmt"


func sender (c chan int){
    c <-1
    c<-2
    c<-3
    c<-4
    close(c)
}
func main(){
    c:=make(chan int ,3)
    go sender (c)
    fmt.Printf("Lenght of the channel c is %v and capacity of channel c is %v",len(c),cap(c))
    for val :=range c{
        fmt.Printf("Length of channel c after value '%v'\n",val,len(c))
    }


}
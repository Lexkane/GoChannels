package main 

import (
    "fmt"
    "math/rand"
    "time"
)

func main (){
    userInput:=make(chan string)
    networkResponses:=make(chan string)
    readUserInput:=delayedResult(userInput)
    makeNetworkRequest:=delayedResult(networkResponses)
    go readUserInput("applesauce")
    go makeNetworkRequest("some data")
    for i:=0;i<2;i++{
        select{
        case userType:=<-userInput:
                fmt.Printf("User typed:%s\n",networkData)
        case networkData:=<-networkResponses:
                fmt.Printf("Data came in from network:%s\n",networkData)
            }
    }
   
}
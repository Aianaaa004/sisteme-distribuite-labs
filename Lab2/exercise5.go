package main

import "fmt"

// Exercise 5: one goroutine sends to another; Receive returns the message;
// main prints the returned value.
func Send(out chan<- string) {
    out <- "mesaj transmis de Send"
}

func Receive(in <-chan string, result chan<- string) {
    message := <-in
    result <- message
}

func main() {
    communication := make(chan string)
    result := make(chan string)

    go Send(communication)
    go Receive(communication, result)

    message := <-result
    fmt.Println("Main a primit:", message)
}

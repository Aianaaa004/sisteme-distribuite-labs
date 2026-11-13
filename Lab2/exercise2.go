package main

import "fmt"

// Exercise 2: main sends a message to another goroutine through a channel.
func main() {
    messages := make(chan string)

    go func() {
        message := <-messages
        fmt.Println("Mesaj primit:", message)
    }()

    messages <- "salut"
}

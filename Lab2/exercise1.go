package main

import "fmt"

// Exercise 1: A goroutine sends a message to main through a channel.
func main() {
    messages := make(chan string)

    go func() {
        messages <- "salut"
    }()

    fmt.Println(<-messages)
}

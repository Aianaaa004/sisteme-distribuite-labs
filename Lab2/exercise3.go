package main

import "fmt"

// Exercise 3: main receives two messages through the same channel.
func main() {
    messages := make(chan string)

    go func() {
        messages <- "primul mesaj"
        messages <- "al doilea mesaj"
    }()

    fmt.Println(<-messages)
    fmt.Println(<-messages)
}

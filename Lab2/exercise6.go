package main

import (
    "fmt"
    "math/rand"
    "time"
)

func robot(name string, finish chan<- string) {
    // Simulate a race. The random delay represents different running times.
    duration := time.Duration(rand.Intn(1500)+500) * time.Millisecond
    time.Sleep(duration)
    finish <- name
}

func main() {
    rand.Seed(time.Now().UnixNano())

    finish := make(chan string, 2)

    fmt.Println("Observer: start! Ambii roboti pornesc simultan.")

    go robot("Robotul 1", finish)
    go robot("Robotul 2", finish)

    // select receives whichever robot finishes first.
    winner := <-finish
    fmt.Println("Observer: castigatorul este", winner)

    // Receive the second result too, so both goroutines can finish cleanly.
    <-finish
}

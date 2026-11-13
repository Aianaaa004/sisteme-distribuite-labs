package main

import "fmt"

// Exercise 4: A goroutine notifies main when its work is complete,
// without using WaitGroup.
func main() {
    finished := make(chan bool)

    go func() {
        fmt.Println("Goroutine: procesez datele...")
        // Simulated work would happen here.
        fmt.Println("Goroutine: am terminat.")
        finished <- true
    }()

    <-finished
    fmt.Println("Main: am primit notificarea de terminare.")
}

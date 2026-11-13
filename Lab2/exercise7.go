package main

import (
    "fmt"
    "sync"
    "time"
)

type Message struct {
    Position [2]int
    Finished bool
}

// A robot sends all its positions to the other robot and then sends
// a final notification. At the same time, it receives the other robot's
// positions until that robot sends its final notification.
func robot(
    name string,
    route [][2]int,
    outgoing chan<- Message,
    incoming <-chan Message,
    observer chan<- string,
    wg *sync.WaitGroup,
) {
    defer wg.Done()

    receiverDone := make(chan struct{})

    go func() {
        defer close(receiverDone)
        for message := range incoming {
            if message.Finished {
                fmt.Println(name, ": celalalt robot a terminat.")
                return
            }
            fmt.Printf("%s a primit pozitia celuilalt robot: (%d, %d)\n",
                name, message.Position[0], message.Position[1])
        }
    }()

    for _, position := range route {
        fmt.Printf("%s se afla la pozitia (%d, %d)\n",
            name, position[0], position[1])
        outgoing <- Message{Position: position}
        time.Sleep(300 * time.Millisecond)
    }

    // Notify the other robot that this route is complete.
    outgoing <- Message{Finished: true}
    <-receiverDone

    observer <- name + " a terminat traseul."
}

func main() {
    // Buffered channels allow both robots to send their route updates
    // while independently receiving the other robot's updates.
    robot1ToRobot2 := make(chan Message, 10)
    robot2ToRobot1 := make(chan Message, 10)
    observer := make(chan string, 2)

    route1 := [][2]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}}
    route2 := [][2]int{{0, 1}, {1, 1}, {1, 2}, {2, 2}}

    var wg sync.WaitGroup
    wg.Add(2)

    go robot("Robotul 1", route1, robot1ToRobot2, robot2ToRobot1, observer, &wg)
    go robot("Robotul 2", route2, robot2ToRobot1, robot1ToRobot2, observer, &wg)

    // Observer receives both completion notifications.
    go func() {
        wg.Wait()
        close(observer)
    }()

    for message := range observer {
        fmt.Println("Observer:", message)
    }

    close(robot1ToRobot2)
    close(robot2ToRobot1)
    fmt.Println("Observer: ambii roboti au terminat.")
}

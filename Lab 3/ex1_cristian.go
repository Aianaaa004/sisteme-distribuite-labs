package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Stocheaza deviatia fata de timpul real
type Clock struct {
	offset time.Duration
}

func (c Clock) Now() time.Time {
	return time.Now().Add(c.offset)
}

// (a) Calculeaza decalajul theta si round-trip-ul delta folosind algoritmul lui Cristian
func estimate(t1, t2, t3, t4 time.Time) (theta, delta time.Duration) {
	// Intarzierea retelei (round-trip delay)
	delta = (t4.Sub(t1)) - (t3.Sub(t2))

	// Offset-ul (decalajul dintre ceasul clientului si cel al serverului)
	theta = (t2.Sub(t1) + t3.Sub(t4)) / 2

	return theta, delta
}

func measure(c_clock, s_clock Clock) (theta, delta time.Duration) {
	// Request (Client -> Server)
	t1 := c_clock.Now()
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	t2 := s_clock.Now()

	// Procesare pe server (intre 5 si 20 ms)
	time.Sleep(5*time.Millisecond + time.Duration(rand.Intn(20))*time.Millisecond)
	t3 := s_clock.Now()

	// Response (Server -> Client)
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	t4 := c_clock.Now()

	return estimate(t1, t2, t3, t4)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Seteaza un ceas client cu un offset initial (de ex. decalat cu +2 secunde)
	clientClock := Clock{offset: 2 * time.Second}
	serverClock := Clock{offset: 0 * time.Second}

	fmt.Printf("Offset initial client: %v\n\n", clientClock.offset)

	// (a) O singura simulare
	theta, delta := measure(clientClock, serverClock)
	fmt.Printf("[Simulare 1] Theta estimat: %v, Delta (RTT): %v\n", theta, delta)

	// Corectarea ceasului clientului pe baza estimarii
	clientClock.offset += theta
	fmt.Printf("Offset client dupa corectie: %v\n\n", clientClock.offset)

	// (b) Repetarea simularii de 10 ori si gasirea valorilor minime
	minTheta := time.Duration(1<<63 - 1)
	minDelta := time.Duration(1<<63 - 1)

	clientClock = Clock{offset: 2 * time.Second} // Resetam offset-ul initial

	for i := 1; i <= 10; i++ {
		th, dl := measure(clientClock, serverClock)
		fmt.Printf("Runda %d -> Theta: %v, Delta: %v\n", i, th, dl)

		if dl < minDelta {
			minDelta = dl
		}
		if th < minTheta {
			minTheta = th
		}
	}

	fmt.Printf("\nValori minime dupa 10 rulari -> Min Theta: %v, Min Delta: %v\n", minTheta, minDelta)
}

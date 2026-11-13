package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type LamportClock struct {
	counter int
}

// 1. Incrementeaza contorul local inainte de un eveniment intern sau trimitere
func (l *LamportClock) Tick() {
	l.counter++
}

// 2. Returneaza valoarea contorului curent pentru a fi atasata ca timestamp la mesaj
func (l *LamportClock) Send() int {
	l.counter++
	return l.counter
}

// 3. Actualizeaza contorul la primirea unui mesaj: max(local, ts) + 1
func (l *LamportClock) Receive(ts int) {
	if ts > l.counter {
		l.counter = ts
	}
	l.counter++
}

type Message struct {
	sender int
	ts     int
}

type Event struct {
	process          int
	internal_counter int
	ts               int
	kind             string // "internal", "send", "recu"
}

// Functie de comparare (less) pentru sortarea evenimentelor dupa regulile lui Lamport
func less(a, b Event) bool {
	if a.ts != b.ts {
		return a.ts < b.ts
	}
	// Daca timestamp-urile sunt egale, se departajeaza dupa ID-ul procesului
	return a.process < b.process
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Initializam ceasurile pentru cele 3 procese
	clocks := map[int]*LamportClock{
		1: {counter: 0},
		2: {counter: 0},
		3: {counter: 0},
	}

	var events []Event

	// Simulare evenimente in inel: P1 -> P2 -> P3 -> P1
	// Eveniment 1: P1 face ceva intern si trimite un mesaj catre P2
	clocks[1].Tick()
	events = append(events, Event{process: 1, ts: clocks[1].counter, kind: "internal"})

	ts1 := clocks[1].Send()
	events = append(events, Event{process: 1, ts: ts1, kind: "send"})

	// P2 primeste mesajul de la P1
	clocks[2].Receive(ts1)
	events = append(events, Event{process: 2, ts: clocks[2].counter, kind: "recu"})

	// Eveniment 2: P2 trimite mesaj catre P3
	ts2 := clocks[2].Send()
	events = append(events, Event{process: 2, ts: ts2, kind: "send"})

	// P3 primeste mesajul de la P2
	clocks[3].Receive(ts2)
	events = append(events, Event{process: 3, ts: clocks[3].counter, kind: "recu"})

	// Eveniment 3: P3 trimite mesaj catre P1
	ts3 := clocks[3].Send()
	events = append(events, Event{process: 3, ts: ts3, kind: "send"})

	// P1 primeste mesajul de la P3
	clocks[1].Receive(ts3)
	events = append(events, Event{process: 1, ts: clocks[1].counter, kind: "recu"})

	// Sortam evenimentele crescator dupa ceasul Lamport folosind functia less
	sort.Slice(events, func(i, j int) bool {
		return less(events[i], events[j])
	})

	// Afisarea rezultatelor ordonate
	fmt.Println("Evenimentele ordonate crescator dupa ceasul Lamport:")
	for _, e := range events {
		fmt.Printf("Proces P%d | Tip: %-8s | Timestamp Lamport: %d\n", e.process, e.kind, e.ts)
	}
}

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Info struct {
	ID    int
	Value int
}

type Memory map[Info]struct{}

type Node struct {
	Neighbors []int
	M         Memory
	Crash     bool
}

func diametru(adj [][]int) int {
	n := len(adj)
	max := 0

	for i := 0; i < n; i++ {
		dist := make([]int, n)

		for j := range dist {
			dist[j] = -1
		}

		dist[i] = 0
		q := []int{i}

		for len(q) > 0 {
			x := q[0]
			q = q[1:]

			for _, v := range adj[x] {
				if dist[v] == -1 {
					dist[v] = dist[x] + 1
					q = append(q, v)
				}
			}
		}

		for _, x := range dist {
			if x == -1 {
				return -1
			}

			if x > max {
				max = x
			}
		}
	}

	return max
}

func main() {
	in := bufio.NewScanner(os.Stdin)

	in.Scan()
	var n, s int
	fmt.Sscan(in.Text(), &n, &s)

	in.Scan()
	valori := strings.Fields(in.Text())

	adj := make([][]int, n)
	nodes := make([]Node, n)

	for i := 0; i < n; i++ {
		in.Scan()

		for _, x := range strings.Fields(in.Text()) {
			v, _ := strconv.Atoi(x)
			adj[i] = append(adj[i], v)
		}
	}

	d := diametru(adj)

	if d == -1 {
		fmt.Println("Graful nu este tare conex")
		return
	}

	for i := 0; i < n; i++ {
		v, _ := strconv.Atoi(valori[i])

		nodes[i].Neighbors = adj[i]
		nodes[i].M = make(Memory)
		nodes[i].M[Info{i, v}] = struct{}{}
	}

	t := 0
	crashuri := 0

	rand.Seed(1)

	for t <= (s+1)*d {

		// Posibil crash
		if crashuri < s && rand.Intn(3) == 0 {
			var candidati []int

			for i := 0; i < n; i++ {
				if !nodes[i].Crash {
					candidati = append(candidati, i)
				}
			}

			if len(candidati) > 0 {
				x := candidati[rand.Intn(len(candidati))]
				nodes[x].Crash = true
				crashuri++

				fmt.Println("Nodul", x, "a intrat in crash")
			}
		}

		// Copiem memoriile
		mem := make([]Memory, n)

		for i := 0; i < n; i++ {
			mem[i] = make(Memory)

			for x := range nodes[i].M {
				mem[i][x] = struct{}{}
			}
		}

		// Trimitem informatiile catre vecini
		for i := 0; i < n; i++ {
			if nodes[i].Crash {
				continue
			}

			for _, v := range nodes[i].Neighbors {
				if nodes[v].Crash {
					continue
				}

				for x := range mem[i] {
					nodes[v].M[x] = struct{}{}
				}
			}
		}

		t++
	}

	fmt.Println("Diametru:", d)
	fmt.Println("Iteratia:", t)
	fmt.Println("Memoria finala:")

	nodFinal := -1

	for i := 0; i < n; i++ {
		if !nodes[i].Crash {
			nodFinal = i
			break
		}
	}

	if nodFinal == -1 {
		fmt.Println("Nu exista noduri functionale")
		return
	}

	var rezultat []Info

	for x := range nodes[nodFinal].M {
		rezultat = append(rezultat, x)
	}

	sort.Slice(rezultat, func(i, j int) bool {
		return rezultat[i].ID < rezultat[j].ID
	})

	for _, x := range rezultat {
		fmt.Printf("(%d, %d) ", x.ID, x.Value)
	}

	fmt.Println()
}

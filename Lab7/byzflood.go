package main

import (
	"fmt"
	"math/rand"
	"time"
)

func majority(values []int) (int, int) {
	zero, one := 0, 0

	for _, v := range values {
		if v == 0 {
			zero++
		} else {
			one++
		}
	}

	if one > zero {
		return 1, one
	}

	return 0, zero
}

func selectByzantine(n, s int) []bool {
	bad := make([]bool, n)

	f := rand.Intn(s + 1)
	indexes := rand.Perm(n)

	for i := 0; i < f; i++ {
		bad[indexes[i]] = true
	}

	return bad
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var n, s int

	fmt.Scan(&n, &s)

	if n <= 4*s || n < s+2 {
		fmt.Println("Invalid input: n must be greater than 4*s")
		return
	}

	values := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&values[i])

		if values[i] != 0 && values[i] != 1 {
			fmt.Println("Values must be 0 or 1")
			return
		}
	}

	bad := selectByzantine(n, s)

	fmt.Println("Nodes:", n)
	fmt.Println("Byzantine faults:")

	for i := 0; i < n; i++ {
		if bad[i] {
			fmt.Println("Node", i, "is Byzantine")
		}
	}

	for t := 0; t < s+2; t++ {
		fmt.Println("Phase:", t)

		// Round 1
		messages := make([][]int, n)

		for i := 0; i < n; i++ {
			messages[i] = make([]int, n)
		}

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if bad[i] {
					messages[j][i] = rand.Intn(2)
				} else {
					messages[j][i] = values[i]
				}
			}
		}

		majorities := make([]int, n)
		counts := make([]int, n)

		for i := 0; i < n; i++ {
			majorities[i], counts[i] = majority(messages[i])
		}

		// Round 2
		leader := t
		fmt.Println("King:", leader)

		tie := make([]int, n)

		if bad[leader] {
			for i := 0; i < n; i++ {
				tie[i] = rand.Intn(2)
			}
		} else {
			for i := 0; i < n; i++ {
				tie[i] = majorities[leader]
			}
		}

		// Update
		newValues := make([]int, n)

		for i := 0; i < n; i++ {
			if float64(counts[i]) > float64(n)/2+float64(s) {
				newValues[i] = majorities[i]
			} else {
				newValues[i] = tie[i]
			}
		}

		values = newValues
	}

	// Final result: correct nodes only
	fmt.Println("Final values of correct nodes:")

	first := -1
	consensus := true

	for i := 0; i < n; i++ {
		if !bad[i] {
			fmt.Println("Node", i, ":", values[i])

			if first == -1 {
				first = values[i]
			} else if values[i] != first {
				consensus = false
			}
		}
	}

	if consensus {
		fmt.Println("Consensus value:", first)
	} else {
		fmt.Println("Consensus not reached")
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 8 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Split(bufio.ScanWords)

	if !scanner.Scan() {
		return
	}
	n, _ := strconv.Atoi(scanner.Text())

	if !scanner.Scan() {
		return
	}
	t, _ := strconv.Atoi(scanner.Text())

	events := make([]int, t+1)
	maxBikes := 0
	currentBikes := 0

	for i := 1; i < n; i++ {
		if !scanner.Scan() {
			break
		}
		a, _ := strconv.Atoi(scanner.Text())

		if !scanner.Scan() {
			break
		}
		f, _ := strconv.Atoi(scanner.Text())

		if !scanner.Scan() {
			break
		}
		s, _ := strconv.Atoi(scanner.Text())

		events[a] += s
		events[f] -= s
	}

	for j := 0; j <= t; j++ {
		currentBikes += events[j]
		if currentBikes > maxBikes {
			maxBikes = currentBikes
		}
	}
	fmt.Println(maxBikes)
}

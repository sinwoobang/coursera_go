package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	fmt.Print("Input integers(ex. 1, 3, 5): ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	var ints []int = make([]int, 0)
	for _, rawInt := range strings.Fields(scanner.Text()) {
		value, _ := strconv.Atoi(rawInt)
		ints = append(ints, value)
	}

	var wg sync.WaitGroup
	var i int

	var unit int = len(ints) / 4
	if unit == 0 {
		unit = 1
	}

	for i = 0; i+unit < len(ints); i += unit {
		wg.Add(1)
		go sortSub(&wg, ints[i:i+unit])
	}
	wg.Wait()

	if i < len(ints) {
		wg.Add(1)
		go sortSub(&wg, ints[i:])
	}
	wg.Wait()

	wg.Add(1)
	go sortSub(&wg, ints)
	wg.Wait()

	fmt.Printf("Final Sorted: %v\n", ints)
}

func sortSub(wg *sync.WaitGroup, ints []int) {
	defer wg.Done()

	sort.Ints(ints)
	fmt.Printf("Sorted: %v\n", ints)
}

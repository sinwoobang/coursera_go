package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var slice []int = make([]int, 0, 3);
	for {
		var input string
		fmt.Scan(&input)
		if input == "X" {
			break;
		}
		i, _ := strconv.Atoi(input)
		slice = append(slice, i)
		sort.Ints(slice)
		fmt.Println(slice)
	}
}
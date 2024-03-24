package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    inputReader := bufio.NewReader(os.Stdin)
    input, _ := inputReader.ReadString('\n')
	input = strings.TrimRight(strings.ToLower(input), "\n");
	t := strings.HasPrefix(input, "i") && strings.Contains(input, "a") && strings.HasSuffix(input, "n")
	if t {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
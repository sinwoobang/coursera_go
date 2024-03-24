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
	input = strings.ToLower(input);
	t := strings.Contains(input, "i") && strings.Contains(input, "a") && strings.Contains(input, "n")
	if t {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
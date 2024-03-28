package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func (t float64) float64 {
		return 1/2 * a * math.Pow(t, 2) + v0 * t + s0;
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin);
	scanner.Scan();
	input_texts := strings.Fields(scanner.Text())
	
	a, _ := strconv.ParseFloat(input_texts[0], 64)
	v0, _ := strconv.ParseFloat(input_texts[1], 64)
	s0, _ := strconv.ParseFloat(input_texts[2], 64)
	
	fn := GenDisplaceFn(a, v0, s0)
	fmt.Println(fn(3))
	fmt.Println(fn(5))
}
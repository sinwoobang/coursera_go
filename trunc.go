package main

import "fmt"

func main() {
	var f float32;

	fmt.Printf("Input float: ");
	fmt.Scan(&f);
	fmt.Printf("%.0f\n", f);
}
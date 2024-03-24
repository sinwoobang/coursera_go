package main

import (
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	var filePath string;

	fmt.Print("Input file path: ")
	fmt.Scan(&filePath)
	
	nameBytes, _ := os.ReadFile(filePath)
	nameString := string(nameBytes)
	lines := strings.Split(nameString, "\n")

	var names []name;
	for _, line := range lines {
		if line == "" {
			continue;
		}

		n := name{fname: strings.Split(line, " ")[0], lname: strings.Split(line, " ")[1]}
		names = append(names, n);
	}

	for _, _name := range names {
		fmt.Printf("first name: %s, last name: %s\n", _name.fname, _name.lname)
	}
}
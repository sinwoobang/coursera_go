package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var name, address string;
	
	fmt.Print("Input name: ");
	if _, err := fmt.Scan(&name); err != nil {
		log.Fatalf("Failed to scan: %v", err);
	}
	
	fmt.Print("Input address: ");
	if _, err := fmt.Scan(&address); err != nil {
		log.Fatalf("Failed to scan: %v", err);
	}

	person := map[string]string{"name": name, "address": address};
	j, _ := json.Marshal(person);
	fmt.Println(string(j));
}
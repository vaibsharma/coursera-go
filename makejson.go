package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var address string

	fmt.Printf("Enter name: ")
	fmt.Scan(&name)
	fmt.Printf("Enter address: ")
	fmt.Scan(&address)

	var m = map[string]string{
		"name":    name,
		"address": address,
	}

	barr, err := json.Marshal(m)
	if err == nil {
		fmt.Print(string(barr))
	} else {
		fmt.Print("Error: ", err)
	}
}

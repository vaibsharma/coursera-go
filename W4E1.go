package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	var name, address string
	var m map[string]string
	m = make(map[string]string)

	fmt.Println("Enter your name: ");
	fmt.Scanf("%s",&name)
	fmt.Println("\nEnter your Address: ")
	fmt.Scanf("%s",&address)
	fmt.Printf("\n");

	m["name"] = name
	m["address"] = address
	j, _ :=  json.Marshal(m)

	fmt.Printf("The byte representation of JSON is : %v\n",j)
	fmt.Printf("The non byte JSON representation is : %v\n",string(j))
}
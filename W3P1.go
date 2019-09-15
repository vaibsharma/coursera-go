package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

func main() {
	var str string
	var a = make([]int, 0, 3)

	for {
		fmt.Scanln(&str)
		num, err := strconv.Atoi(str)

		if str == "X" {
			break
		}
		if err != nil{
			log.Fatal(err)
		}
		
		a = append(a, num)
		sort.Ints(a)
		fmt.Printf("The new sorted array after inserting %d is %v\n",num, a)
	}
}

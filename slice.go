package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	s := make([]int, 0, 3)

	var i string
	for {
		fmt.Print("Enter an integer: ")
		_, err := fmt.Scan(&i)
		if err != nil || i == "X" {
			break
		} else {
			ii, err := strconv.Atoi(i)
			if err != nil {
				fmt.Println("Enter an integer or 'X'!")
			} else {
				s = append(s, ii)
				printSortedSlice(s)
			}
		}
	}
}

func printSortedSlice(s []int) {
	sort.Ints(s)
	fmt.Printf("%v\n", s)
}

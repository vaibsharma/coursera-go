package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	s := make([]name, 0, 1)

	var fileName string
	fmt.Print("Enter file name: ")
	fmt.Scan(&fileName)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		temp := strings.Fields(line)
		s = append(s, name{temp[0], temp[1]})
	}
	file.Close()
	printSlice(s)
}

func printSlice(s []name) {
	for _, emt := range s {
		fmt.Println(emt.fname, emt.lname)
	}
}

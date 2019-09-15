package main

import (
	"fmt"
	// "encoding/json"
	"log"
	"bufio"
	"os"
	"strings"
	// "io"
)

type Name struct{
	Fname string
	Lname string
}

func main(){
	file, err := os.Open("go.text")
	var ans []Name
	if err != nil {
		log.Fatal(err)
	}
	
	defer file.Close()

	reader := bufio.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println(err)
			break
		}
		a := strings.Split(string(line), " ")
		ans = append(ans, Name{Fname: a[0], Lname: a[1]})
	}
	for ind, obj := range(ans){
		fmt.Printf("%d. First Name: %s, Last Name: %s \n", ind, obj.Fname, obj.Lname)
	}
}
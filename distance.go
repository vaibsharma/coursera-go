package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
	"strconv"
)

func main() {

	//Get Datas acceleration, initial velocity, and initial displacement
	a  := getInput("acceleration")
	v0 := getInput("initial velocity")
	i0 := getInput("initial displacement")
	fmt.Printf("a= %f , v0= %f , i0= %f \n",a,v0,i0)

	disFunc := GenDisplaceFn(a,v0,i0)
	//Get time
	t := getInput("time in second")
	fmt.Printf("waite time:%f \n",t)

	fmt.Printf("\nGenDisplaceFn:%f",disFunc(t))
}

func GenDisplaceFn(a,v0,i0 float64) func(float64) float64 {
	//func to calculate
	fn:= func(t float64) float64 {
		//duration := time.Duration(int(t))*time.Second
		//fmt.Printf("wait t:%dsec to calculate GenDisplaceFn ",int(t))
		// wait to calculate
		//time.Sleep( duration)
		return 0.5 * a * t * t + v0 * t + i0
	}
	 return fn
}


func getInput(ikey string  ) (entered float64) {
	reader := bufio.NewReader(os.Stdin)
	isOk:= true

	for isOk {
		fmt.Print("Enter ",ikey,":")
		s, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("err!",err)
		}else{
			s = strings.TrimSuffix(s, "\n")
			i, err2 := strconv.ParseFloat(s,64)
			if err2 != nil {
				fmt.Println("invalid number!!!", err2)
			} else {
				isOk = false
				entered= i
			}
		}
	}

	return entered
}
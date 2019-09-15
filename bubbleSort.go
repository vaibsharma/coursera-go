package main

import (
	"fmt"
)

func swap(arr []int, i int) {
	allias := arr[i]
	arr[i] = arr[i+1]
	arr[i+1] = allias
}

func BubbleSort (arr []int){
	num := len(arr)
	for x := 0; x < num; x++{
		for y:=0;y < num - x - 1; y++{
			if arr[y] > arr[y+1]{
				swap(arr, y)
			}
		}
	}
}

func main(){
	var arr[]int
	var num int
	fmt.Println("Please enter 10 numbers to be sorted: ");
	fmt.Scanf("%d",&num)
	for x,y := 0,1; x < 10 ; x++ {
		fmt.Scanf("%d",&y)
		arr = append(arr, y)
	}
	BubbleSort(arr)
	fmt.Printf("\nThe sorted array after bubblesort is %v",arr)
}

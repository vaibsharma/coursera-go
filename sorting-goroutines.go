package main

import (
	"fmt"
	"math"
	"sort"
	"sync"
)

func main() {
	var arrLen uint64
	arr := []uint64{}
	var wg sync.WaitGroup
	fmt.Printf("Please enter the length of array: ")
	fmt.Scanf("%d", &arrLen)

	for x, num := 0, 0; x < int(arrLen); x++ {
		_, err := fmt.Scanf("%d", &num)
		if err != nil {
			fmt.Printf("Trust issues :(")
		}
		arr = append(arr, uint64(num))
	}

	for i := 0; i < 4 && i*4 < int(arrLen); i++ {
		from := i * 4
		to := math.Min(float64(i+1)*4, float64(arrLen))

		wg.Add(1)
		// fmt.Printf("Sorting from %d to %d ", int(from), int(to))
		go func(arr []uint64) {
			sort.Slice(arr, func(p, q int) bool {
				return arr[p] < arr[q]
			})
			fmt.Printf("The sorted array from index %d to index %d is %v \n", int(from), int(to), arr)
			wg.Done()
		}(arr[int(from):int(to)])

		wg.Wait()
	}
	fmt.Printf("After the goroutines the array is %v \n", arr)
	sort.Slice(arr, func(p, q int) bool {
		return arr[p] < arr[q]
	})
	fmt.Printf("After final sorting the array is %v \n", arr)
}

package main 

import (
		"fmt"
)

func bubbleSort(arr []int) []int {
	swapped := true 
	for swapped {
		swapped = false
		for i := 0;  i < len(arr) - 1; i++ {

			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
			fmt.Println("Status : " , arr)
		}
	}
	return arr
}

func main() {
	var myarr, myar []int
	myarr = []int{9, 5, 3, 12, 42, 22, 18, 27, 44}
	myar = bubbleSort(myarr)
	fmt.Println(myar)
}

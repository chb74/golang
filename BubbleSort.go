// -- By liverpools@gmail.com 
// -- On Feb 5, 2022 
package main 

import ( "fmt" )

func BubbleSort(arr []int) []int {
	fmt.Println(arr)
	for i := 0; i < len(arr) - 1; i++ {
		for j := 0; j < len(arr) - 1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	var arr, result []int 

	arr = []int{ 10, 9, 20, 5, 3, 1, 25}

	result = BubbleSort(arr)
	fmt.Println("Result is " , result)
}
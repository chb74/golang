// -- By liverpools@gmail.com 
// -- On Feb 5, 2022 

package main 

import ( "fmt" )

func SelectionSort(arr []int) []int {

	for i := 0; i < len(arr) - 1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}

	return arr
}

func main() {

	var myarr, result []int 
	myarr = []int{10, 8, 5, 2, 1, 4, 7}

	result = SelectionSort(myarr)
	fmt.Println("Result is ", result)
}
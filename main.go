package main

import "fmt"

func sort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		a := i - 1
		for a >= 0 && arr[a] > key {
			arr[a+1] = arr[a]
			a--
		}
		arr[a+1] = key
	}
	return arr
}

func main() {
	array := []int{5,2,6,1,3,7}
	sortedArray := sort(array)
	fmt.Printf("Sorted Array: %v\n", sortedArray)
}

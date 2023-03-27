package main

import "fmt"

func binarySearch(serching int, array []int) bool {

	low := 0
	high := len(array) - 1

	for low <= high {
		median := (low + high) / 2

		if array[median] < serching {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(array) || array[low] != serching {
		return false
	}

	return true
}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	var num int
	fmt.Println("Enter number for binary search\ntrue == find\nfalse == not find ")
	fmt.Scan(&num)
	fmt.Println(binarySearch(num, items))
}

package main

import (
	"fmt"
)

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	center := len(a) / 2

	a[center], a[right] = a[right], a[center]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]

	quickSort(a[:left])
	quickSort(a[left+1:])
	return a
}

func main() {
	arr := []int{1, 10, 30, 5, 2, 11, 12, 6, 54, 123}
	quickSort(arr)

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

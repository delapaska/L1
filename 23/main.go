package main

import "fmt"

func remove(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)

}

func main() {
	arr := []int{10, 20, 34, 42, 51, 66, 74, 83, 92, 108}

	fmt.Println("Your Array: ", arr)

	fmt.Println("Enter index element that u want to delete:")
	var x int
	fmt.Scan(&x)
	arr = remove(arr, x)
	fmt.Println(arr)
}

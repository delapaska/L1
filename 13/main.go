package main

import "fmt"

func main() {
	var x int
	var y int
	fmt.Println("Enter x and y:")
	fmt.Scan(&x)
	fmt.Scan(&y)
	fmt.Printf("x:%d  y:%d\n", x, y)
	//1
	x = x + y
	y = x - y
	x = x - y

	//2
	//x,y = y,x

	fmt.Printf("Processing...\nx:%d  y:%d", x, y)

}

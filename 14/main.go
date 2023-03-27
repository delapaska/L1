package main

import "fmt"

func main() {
	var i int
	fmt.Println(getType(i))

	var b bool
	fmt.Println(getType(b))

	var s string
	fmt.Println(getType(s))

	var f32 float32
	fmt.Println(getType(f32))

	var f64 float64
	fmt.Println(getType(f64))
}

func getType(v interface{}) string {

	return fmt.Sprintf("It's: %T", v)
}

package main

import "fmt"

func readFloat32(x *float32) {
	var y float32 = 0
	fmt.Scan(&y)
	*x = y
}

func readInt(x *int) {
	var y int = 0
	fmt.Scan(&y)
	*x = y
}

func main() {

	// Ввод размера массива и его содержимого
	var n int = 0
	fmt.Println("Введите количество элементов:")
	readInt(&n)
	arr := make([]float32, n)
	fmt.Println("Введите элементы:")
	var k float32 = 0
	for i := 0; i < n; i++ {
		readFloat32(&k)
		arr[i] = k
	}
	grouping := make(map[int][]float32)

	for _, elements := range arr {
		key := (int(elements) / 10) * 10
		grouping[key] = append(grouping[key], elements)
	}

	fmt.Println(grouping)

}

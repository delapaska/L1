package main

import "fmt"

func intersection(set1, set2 []int) []int {
	// создаем map для хранения элементов первого множества
	// инициализируем его значениями из первого множества
	tempMap := make(map[int]bool)
	for _, val := range set1 {
		tempMap[val] = true
	}

	// создаем результирующий slice
	var result []int

	// проверяем, есть ли элементы из второго множества
	// в map первого множества
	for _, val := range set2 {
		if tempMap[val] {
			result = append(result, val)
		}
	}

	return result
}

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}
	res := intersection(set1, set2)
	fmt.Println(res) // [3 4 5]
}

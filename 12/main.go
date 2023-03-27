package main

import "fmt"

//доработать
func main() {
	var strings = []string{"cat", "name", "cat", "dog", "cat", "tree", "name"}
	fmt.Println(createProperSet(strings))
}

func createProperSet(strings []string) []string {
	m := make(map[string]bool)
	for _, item := range strings {
		if _, ok := m[item]; ok {
			// вывод повторяющихся значений
			// duplicate item
			//fmt.Println(item, "is a duplicate")
		} else {
			m[item] = true
		}
	}

	var result []string
	for item, _ := range m {
		result = append(result, item)
	}
	return result

}

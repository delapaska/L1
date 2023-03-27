package main

import (
	"fmt"
	"strconv"
)

// Создаём 2 структуры
type Human struct {
	name       string
	surname    string
	birth_year int
}

type Action struct {
	Human
}

// Функции по созданию экземпляра и взаимодействия с полями структуры
func Create(name string, surname string, birth_year int) Human {
	hum := Human{
		name:       name,
		surname:    surname,
		birth_year: birth_year,
	}
	return hum
}
func (h *Human) TakeAge() {
	age := 2023 - h.birth_year
	str_age := strconv.Itoa(age)
	fmt.Printf("Human age is %s\n", str_age)
}

func (h *Human) printData() {
	fmt.Printf("His name is %s and surname is %s\n", h.name, h.surname)
}

func main() {
	Hum := Create("Ivan", "Ivanov", 2003)
	Act := Action{Human: Hum}
	Act.TakeAge()
	Act.printData()

}

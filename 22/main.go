package main

import (
	"fmt"
	"math/big"
)

func main() {
	var example float64
	fmt.Println("Enter a and b")
	fmt.Scan(&example)
	a := big.NewFloat(example)
	fmt.Scan(&example)
	b := big.NewFloat(example)
	var input string
	c := big.NewFloat(0.0)

	for {
		fmt.Println("Enter one: * / + - or exit")
		fmt.Scan(&input)
		switch input {
		case "*":
			fmt.Printf("a * b = %f\n", c.Mul(a, b))
			break

		case "/":
			fmt.Printf("a / b = %f\n", c.Quo(a, b))
			break
		case "+":
			fmt.Printf("a + b = %f\n", c.Add(a, b))
			break
		case "-":
			fmt.Printf("a - b = %f\n", c.Sub(a, b))
			break
		case "exit":
			fmt.Println("Bye!")
			return

		}
	}

}

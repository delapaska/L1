package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

// Конструктор
func CreatePoint(x int, y int) *Point {
	return &Point{x, y}
}

// Вычисление расстояния
func CalcDistanc(point1 *Point, point2 *Point) float64 {
	x := point1.x - point2.x
	y := point1.y - point2.y
	dist := math.Sqrt(float64(x*x + y*y))
	return dist

}

func main() {
	point_1 := CreatePoint(5, 5)
	point_2 := CreatePoint(1, 2)
	ans := CalcDistanc(point_1, point_2)

	fmt.Println(ans)
}

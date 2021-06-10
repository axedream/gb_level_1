package main

import (
	"fmt"
	"math"
)

func main() {

	var s float64

	fmt.Println("Программа вычисляющая диаметр и длину окружности по заданной площади круга")
	fmt.Print("Введите площадь круга: ")
	fmt.Scanln(&s)

	r := math.Sqrt(s/math.Pi)

	//диаметр окружности
	d := 2*r

	fmt.Println("\nДиаметр окружности: ",d)

	//длиннна окружности
	l := 2*math.Pi*r

	fmt.Println("\nДлинна окружности: ",l)

}
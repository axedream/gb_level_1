package main

import (
	"fmt"
)

func main() {
	var x,y int
	fmt.Println("Программа вычисления площади прямоугольника")
	fmt.Print("Введите длину первой стороны прямоугольника X: ")
	fmt.Scanln(&x)
	fmt.Print("Введите длину второй строны приямоугольника Y: ")
	fmt.Scanln(&y)

	result := x * y

	fmt.Printf("Площадь прямоугольника %d",result)

}

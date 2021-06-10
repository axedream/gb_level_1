package main

import (
	"fmt"
	"math"
)

func main()  {
	var op1,op2,result float64
	var op,textOper string
	var textOperKey int


		fmt.Println("Введите первый операнд (только число): ")
	if _,err := fmt.Scanln(&op1); err != nil {
		fmt.Println("Неверно введен операнд 1")
		return
	}
	fmt.Println("Введите операцию (+-*/): ")
	if _,err := fmt.Scanln(&op); err !=nil {
		fmt.Println("Неверно введена операция")
		return
	}
	fmt.Println("Введите второй операнд (только число): ")
	if _,err := fmt.Scanln(&op2); err != nil {
		fmt.Println("Неверно введен операнд 2")
		return
	}

	//пол умолчанию 0 - ошибки нет, 1 - ошибка операнда, 2 - ошибка деления на 0
	textOperKey = 0

	switch op {
	case "+":
		result = op1 + op2
		textOper = "сложение"

	case "-":
		result = op1 - op2
		textOper = "вычитание"

	case "*":
		result = op1 * op2
		textOper = "умножение"

	case "/":
		result = op1 / op2;
		if op2 == 0 { textOperKey = 2 } else {
			textOper = "деление"
		}
	case "^":
		result = math.Pow(op1, op2)
		textOper = "возведение в степень"

	case "per":
		if op1==0 || op2 == 0 {
			textOperKey = 3
		}
		result = 2* (op1 + op2)
		textOper = "периметр прямоугольника"



	default:
		textOperKey = 1
	}

	switch textOperKey {
	case 1 :
		fmt.Println( "Ошибка. Операция не распознана!")
	case 2:
		fmt.Println( "Ошибка. В операции деление, нельзя делить на 0")
	case 3:
		fmt.Println("Ошибка. При вычислении периметра не должно быть сторон равных 0")
	default:
		fmt.Printf("Операнд 1: [%f64] Операция: [%s] операнд 2: [%f64] результат = [%f64]", op1,textOper,op2,result)
	}

}

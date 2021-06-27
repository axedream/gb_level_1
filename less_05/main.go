package main

import (
	"fmt"
)
/**
	Метод рекурсии
 */
func FibonacchiRec(n int) int{

	if n == 0 {
		return 0;
	}

	if n == 1 {
		return 1;
	}

	return FibonacchiRec(n-2) + FibonacchiRec(n-1);
}
/**
	Сохраняем предыдущий результат в маппе
 */


func FibonacchiMap(n int) int{
	Mapa := map[int]int {
		1 : 1,
		2 : 1,
	}

	if n == 1 || n==2 {
		return Mapa[n]
	}

	for i:=3; i<=n; i++ {
		Mapa[i] = Mapa[i-1] + Mapa[i-2]
		//fmt.Println("Значение i:",i, "Значение Mapa", Mapa)
	}

	return Mapa[n]
}


func main()  {
	var opf int

	fmt.Println("Введите число для растчета числа Фибоначчи (положительное, целое >0 ): ")
	if _,err := fmt.Scanln(&opf); err != nil || opf ==0{
		fmt.Println("Введены не верные данные!")
		return
	}

	fmt.Println("Число Фибоначчи (расчитанное рекурсией): ",FibonacchiRec(opf))


	fmt.Println("Число Фибоначчи (расчитанное оптимизированная с помощью Мапы): ",FibonacchiMap(opf))

}

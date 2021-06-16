package main

import (
	"fmt"
)

func sortIntSlice(data[]int) []int {
	//перебираем текущий срез
	for i := 1; i < len(data); i++ {
		//устанавливаем текущий элемент
		nowData := data[i]
		//устанавливаем текущий индекс
		j := i

		for ; j >= 1 && data[j-1] > nowData; j-- {
			data[j] = data[j-1]
		}

		//присваиваем в новое место
		data[j] = nowData
	}

	return data
}


func main()  {
	var data = []int{50,6,8,12,56,34,5}
	fmt.Println("Сортировка среза вставками! Начальный срез: ", data)
	fmt.Println("Отсортированный срез: ", sortIntSlice(data))


}

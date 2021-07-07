package main

import (
	sorting2 "axe_dream/sort/sorting"
	"fmt"
)

func main()  {

	//функция сортировки
	var data = []int{50,6,8,12,56,34,5}
	fmt.Println("Сортировка среза вставками! Начальный срез: ", data)
	sorting := sorting2.SortIntSlice(data)
	fmt.Println("Отсортированный срез: ", sorting)

}


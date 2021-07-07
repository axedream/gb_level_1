package sorting

func SortIntSlice(data[]int) []int {
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
package main

import (
	"fmt"
	"strings"
)

func main() {

	var strInp string

	fmt.Println("Программа определяющая в трехзначном числе количество сотен, десятков и единиц")
	fmt.Print("Введите трехзначное число: ")
	fmt.Scanln(&strInp)

	s := strings.Split(strInp, "")

	for index, item := range s {
		switch index {
		case 0:
			fmt.Println("Число сотен:", item)
		case 1:
			fmt.Println("Число десятков:", item)
		case 2:
			fmt.Println("Число единиц:", item)
		}
	}


}
package main

import (
	"./cnf"
	"./server"
)

type Config struct {
	Port string
	Name string
	Count int
}

func main()  {
	c := cnf.NewConfig()

	//загрузка конфигурации из переменных окружения
	//c.LoadConfigEnv()

	//показываем что загрузили
	//fmt.Println(c)

	//загрузка конфигурации из файла yaml
	//c.LoadConfigFileYaml()

	//загрузка конфигурации из файла json
	//c.LoadConfigFileJson()

	//автоопределения файла
	c.LoadConfiFile()


	//стартуем сервер с нужной конфигурацией
	server.StartServer(c)
}

//для установки требуется установить пакет: go get gopkg.in/yaml.v2
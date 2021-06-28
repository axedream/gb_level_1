package main

import (
	"fmt"
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
	c.LoadConfig()
	fmt.Println(c)
	server.StartServer(c)
}
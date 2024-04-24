package main

//go:generate go build -tags=gen
//go:generate ./using_go_generate.exe

import (
	"fmt"
)

var greetings string = "Программа была запущена стандартным способом"

func main() {
	fmt.Println(greetings)
}

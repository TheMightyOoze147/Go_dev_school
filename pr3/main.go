package main

import (
	"flag"
	"fmt"
	"os"
	"pr3/calculations"
	"strconv"
)

func main() {
	// Получение числа n из аргументов командной строки и преобразование из строки в число
	nvalues := os.Args[1:]
	n_number, _ := strconv.Atoi(nvalues[1])
	// Получение значения флага -log
	var log = flag.Bool("log", false, "true - вывод лога, false - без вывода")
	flag.Parse()
	val := calculations.Calculate(int64(n_number), *log)
	fmt.Println("Answer is:", val)
}

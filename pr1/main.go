package main

import (
	"errors"
	"fmt"
)

func plus(x, y float64) float64 {
	return x + y
}

func minus(x, y float64) float64 {
	return x - y
}

func multi(x, y float64) float64 {
	return x * y
}

func divi(x, y float64) float64 {
	var err = errors.New("деление на ноль")
	if y == 0 {
		fmt.Println("Возникла ошибка: ", err)
		return 0
	} else {
		return x / y
	}

}

func main() {
	var val1, val2 float64
	var choise string
	loop := true
	for loop {
		fmt.Println("\nНапиши опцию и два числа: \n+ - Просуммировать\n- Отнять\n* - Умножить\n/ - Поделить\n^C- Выйти")
		fmt.Scan(&choise, &val1, &val2)
		switch choise {
		case "+":
			fmt.Println("Ответ: ", plus(val1, val2))
		case "-":
			fmt.Println("Ответ: ", minus(val1, val2))
		case "*":
			fmt.Println("Ответ: ", multi(val1, val2))
		case "/":
			fmt.Println("Ответ: ", divi(val1, val2))
		default:
			fmt.Println("Неправильно введена операция или значения")
		}
	}
}

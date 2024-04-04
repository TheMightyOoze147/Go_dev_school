package main

import (
	"errors"
	"fmt"
	"slices"
)

func main() {
	// Создание слайса и добавление в него чисел
	var array_length, to_append, k_index int
	var sli []int
	var len_err = errors.New("значение k превышает длину массива")
	fmt.Println("Длина массива | Значения | k-е максимальное")
	fmt.Scan(&array_length)
	for i := 0; i < array_length; i++ {
		fmt.Scan(&to_append)
		sli = append(sli, to_append)
	}

	fmt.Scan(&k_index)

	// Удаление максимального числа из слайса, вывод последнего "выжившего" максимума
	max_value := 0
	if k_index > array_length {
		fmt.Println("Возникла ошибка:", len_err)
	} else {
		for i := 0; i < k_index; i++ {
			max_value = slices.Max(sli)
			for j, val := range sli {
				if val == max_value {
					sli = append(sli[:j], sli[j+1:]...)
				}
			}
		}
		fmt.Println(max_value)
	}
}

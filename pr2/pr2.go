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
	// Создание k слайсов для поиска в них максимальных значений
	// max_slices := make([][]int, k_index)
	// if k_index > array_length {
	// 	fmt.Println("Возникла ошибка: ", len_err)
	// } else {
	// 	start_index := 0
	// 	for i := 0; i < k_index; i++ {
	// 		part_length := array_length / k_index
	// 		if i+1 == k_index && array_length%k_index != 0 {
	// 			if array_length-(i*(array_length/k_index)+1) == k_index {
	// 				part_length = k_index
	// 			} else {
	// 				part_length = array_length%k_index + k_index
	// 			}
	// 		}
	// 		end_index := start_index + part_length
	// 		max_slices[i] = sli[start_index:end_index]
	// 		start_index = end_index
	// 	}
	// }
	// fmt.Println(max_slices)

	// k_max := 0
	// var max_values []int
	// for i := 0; i < k_index; i++ {
	// 	max_values = append(max_values, slices.Max(max_slices[i]))
	// }
	// k_max = slices.Min(max_values)
	// fmt.Println(max_values)
	// fmt.Println(k_max)

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

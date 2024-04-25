package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func people_counter(amount int, counted *int, number *int) {
	// Создаём мьютекс чтобы предупредить состояние гонки в дальнейшем
	var mutex sync.Mutex
	for i := 0; i < amount; i++ {
		mutex.Lock()
		*number++
		*counted++
		mutex.Unlock()
		time.Sleep(time.Duration(rand.Intn(3)+2) * time.Second) // Время прохождения рамки 2 - 5 сек
	}
}

func frame(amount int) {
	var counted_people int = 0
	var counted_first_stream int = 0
	var counted_second_stream int = 0

	// Поток "людей" делится на два, по количеству рамок
	firstStream := amount / 2
	secondStream := amount - firstStream // Подходит даже для нечётного случая

	// Запуск горутин
	go people_counter(firstStream, &counted_people, &counted_first_stream)
	go people_counter(secondStream, &counted_people, &counted_second_stream)

	// Вывод работы горутин (изменения переменных)
	for {
		fmt.Printf("\n-------Всего------\n-------- %d --------\n-Первая-----Вторая-\n--- %d ------- %d ---\n",
			counted_people, counted_first_stream, counted_second_stream)
		if counted_people == amount {
			break
		}
		time.Sleep(time.Second)
	}
}

func main() {
	frame(21)
	fmt.Println("Перерыв на обед.")
}

# Задание №6: Многопоточное программирование
## Цель: Применить на практике полученные знания многопоточного программирования в языке Go. Использовать примитивы синхронизации доступа к общим разделяемым ресурсам.
## Задание:
В данной работе предстоит выполнить следующие задачи:
1. Реализовать структуру с использованием примитивов синхронизации Go.
2. Продемонстрировать работу созданной структуры с использованием выбранного примитива синхронизации.
3. Продемонстрировать работу программы без использования примитива синхронизации. Предполагается, что программа может работать только с его использованием. Например: Написать структуру счетчика людей, проходящих через турникет в метро. В программе должна использоваться многопоточность.

## Выполнение задания
Написана программа, эмулирующая работу двух рамок для подсчёта проходящих через одну и обе рамки людей.
### С использованием примитива синхронизации
Работа с переменными синхронизирована с помощью mutex.Lock()
```go
mutex.Lock()
*number++
*counted++
mutex.Unlock()
```
Люди проходят рамку от 2-х до 5-ти секунд
```go
time.Sleep(time.Duration(rand.Intn(3)+2) * time.Second)
```
Работа программы при запуске для 21 одного человека функцией frame(21)
<details>

    >go run main.go
    -------Всего------
    -------- 0 --------
    -Первая-----Вторая-
    --- 0 ------- 0 ---

    -------Всего------
    -------- 2 --------
    -Первая-----Вторая-
    --- 1 ------- 1 ---

    -------Всего------
    -------- 3 --------
    -Первая-----Вторая-
    --- 1 ------- 2 ---

    -------Всего------
    -------- 4 --------
    -Первая-----Вторая-
    --- 2 ------- 2 ---

    -------Всего------
    -------- 6 --------
    -Первая-----Вторая-
    --- 3 ------- 3 ---

    -------Всего------
    -------- 8 --------
    -Первая-----Вторая-
    --- 4 ------- 4 ---

    -------Всего------
    -------- 9 --------
    -Первая-----Вторая-
    --- 5 ------- 4 ---
    
    ...
    ...
    ...

    -------Всего------
    -------- 21 --------
    -Первая-----Вторая-
    --- 10 ------- 11 ---
    Перерыв на обед.
</details>

### Без использования примитива синхронизации
Если убрать mutex.Lock(), то редко, но случаются моменты неправильного подсчёта общего количества человек, прошедших обе рамки из-за гонки за ресурс двух горутин. Из-за реализации в программе выхода из цикла для вывода информации, когда счётчик дойдёт до 21, выход из цикла происходит тогда, когда общее число людей в переменной counted_people станет равно 21, но сумма счётчиков двух рамок останется меньше общего количества людей.
<details>

    >go run main.go
    -------Всего------
    -------- 0 --------
    -Первая-----Вторая-
    --- 0 ------- 0 ---

    -------Всего------
    -------- 2 --------
    -Первая-----Вторая-
    --- 1 ------- 1 ---

    -------Всего------
    -------- 3 --------
    -Первая-----Вторая-
    --- 1 ------- 2 ---

    -------Всего------
    -------- 4 --------
    -Первая-----Вторая-
    --- 2 ------- 2 ---

    -------Всего------
    -------- 6 --------
    -Первая-----Вторая-
    --- 2 ------- 3 ---

    -------Всего------
    -------- 7 --------
    -Первая-----Вторая-
    --- 3 ------- 3 ---

    ...
    ...
    ...

    -------Всего------
    -------- 21 --------
    -Первая-----Вторая-
    --- 9 ------- 11 ---
    Перерыв на обед.
</details>
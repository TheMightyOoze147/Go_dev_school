# Использование go:generate
Реализована программа, которая использует go:generate для сборки с флагом gen
```go
//go:generate go build -tags=gen
//go:generate ./using_go_generate.exe
```
Запуск с go generate
```console
>go generate
Программа была запущена с помощью go:generate!
```
Запуск с помощью go run
```console
>go run main.go
Программа была запущена стандартным способом
```
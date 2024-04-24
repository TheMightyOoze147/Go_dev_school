# Использование директивы go:embed
Реализована программа с использованием директивы go:embed, которая выводит содержимое файла text_file.txt

```go
//go:embed text_file.txt
var readme string
```

```console
>go run main.go
Этот текст будет отображаться в консоли благодаря директиве go:embed
```
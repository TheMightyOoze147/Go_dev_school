# Сборка при помощи тэгов
Реализована программа с двумя тэгами ru и en для отображения сообщения на выбранном языке в консоли
## Тэг ru:
```console
>go build -tags=ru
>./uilding_w_tags.exe
Сейчас выбран русский язык!
```
## Тэг en:
```console
>go build -tags=en
>./uilding_w_tags.exe
Current language is english!
```
## Сборка без флагов: 
```console
>go build
>./uilding_w_tags.exe
English language is using as default
```
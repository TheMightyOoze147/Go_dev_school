# Задание №5: Тестирование и линтинг в Go

## Цель: Покрыть заранее предоставленный код тестами на уровне 80% покрытия, использовать mock-тесты для проверки различных сценариев, а также выполнить анализ кода с помощью golangci-lint и исправить выявленные замечания.

## Задание:
В данной работе предстоит выполнить следующие задачи:

1. Тестирование: Написать тесты для заранее предоставленного кода с целью достижения покрытия минимум 80%.
2. Линтинг с помощью golangci-lint: Запустить golangci-lint на предоставленном коде для выявления потенциальных проблем, стилевых нарушений и ошибок. Исправить все выявленные замечания.

## Выполнение задания
### Тесты
Разработаны тесты для пакетов wifi и db с покрытием более 80%:
```console
>go test -cover internal/db
coverage: 80.6% of statements
>go test -cover internal/wifi
coverage: 82.3% of statements
```
Для тестирования использовались mock, эмулирующие работу базы данных и wifi
### Линтер
Вывод линтера до исправления ошибок
<details>
        >golangci-lint run      
        level=warning msg="[linters_context] copyloopvar: this linter is disabled because the Go version (1.19) of your project is lower than Go 1.22"
        level=warning msg="[linters_context] intrange: this linter is disabled because the Go version (1.19) of your project is lower than Go 1.22"
        internal\db\db_functions.go:34:4: exitAfterDefer: log.Fatal will exit, and `defer rows.Close()` will not run (gocritic)
                                log.Fatal(err)
                                ^
        internal\db\db_functions.go:59:4: exitAfterDefer: log.Fatal will exit, and `defer rows.Close()` will not run (gocritic)
                                log.Fatal(err)
                                ^
        internal\db\db_functions.go:12:6: exported: type name will be used as db.DBService by other packages, and that stutters; consider calling this Service (revive)
        type DBService struct {
        ^
        internal\db\db_functions.go:25:15: error returned from interface method should be wrapped: sig: func (example_mock/internal/db.Database).Query(query string, args ...any) (*database/sql.Rows, error) (wrapcheck)
                        return nil, err
                                ^
        internal\db\db_functions.go:40:15: error returned from external package is unwrapped: sig: func (*database/sql.Rows).Err() error (wrapcheck)
                        return nil, err
                                ^
        internal\db\db_functions.go:43:16: error returned from interface method should be wrapped: sig: func (example_mock/internal/db.Database).Query(query string, args ...any) (*database/sql.Rows, error) (wrapcheck)
                return names, err
                        ^
        internal\db\db_functions.go:50:15: error returned from interface method should be wrapped: sig: func (example_mock/internal/db.Database).Query(query string, args ...any) (*database/sql.Rows, error) (wrapcheck)
                        return nil, err
                                ^
        internal\db\db_functions.go:65:15: error returned from external package is unwrapped: sig: func (*database/sql.Rows).Err() error (wrapcheck)
                        return nil, err
                                ^
        internal\db\db_functions.go:52:2: only one cuddle assignment allowed before defer statement (wsl)
                defer rows.Close()
                ^
        internal\db\db_functions.go:61:3: append only allowed to cuddle with appended value (wsl)
                        values = append(values, value)
                        ^
        internal\db\db_functions.go:36:3: append only allowed to cuddle with appended value (wsl)
                        names = append(names, name)
                        ^
        internal\db\db_functions.go:49:2: only one cuddle assignment allowed before if statement (wsl)
                if err != nil {
                ^
        internal\wifi\wi-fi.go:26:2: Consider pre-allocating `addrs` (prealloc)
                ^
        internal\wifi\wi-fi.go:40:2: Consider pre-allocating `name_list` (prealloc)
                var name_list []string
                ^
        internal\wifi\wi-fi.go:13:6: exported: type name will be used as wifi.WiFiService by other packages, and that stutters; consider calling this Service (revive)
        type WiFiService struct {
        ^
        internal\wifi\wi-fi.go:45:2: return with no blank line before (nlreturn)
                return name_list, nil
                ^
        internal\wifi\wi-fi.go:24:15: error returned from interface method should be wrapped: sig: func (example_mock/internal/wifi.WiFi).Interfaces() ([]*github.com/mdlayher/wifi.Interface, error) (wrapcheck)
                        return nil, err
                                ^
        internal\wifi\wi-fi.go:38:15: error returned from interface method should be wrapped: sig: func (example_mock/internal/wifi.WiFi).Interfaces() ([]*github.com/mdlayher/wifi.Interface, error) (wrapcheck)
                        return nil, err
                                ^
        cmd\wifi\main.go:15:3: return with no blank line before (nlreturn)
                        return
                        ^
        cmd\wifi\main.go:23:3: return with no blank line before (nlreturn)
                        return
                        ^
</details>

Все ошибки в коде были исправлены, линтер более не выводит никаких ошибок.
<details>
<pre>
P.S. После линтера тесты сломались
</pre>
<img src="https://media.tenor.com/5aF7np_zPEgAAAAe/pepe-why-pepe-the-frog.png" width="20%" height="20%"/>
</details>
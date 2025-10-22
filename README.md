Server Checker

Простая утилита для периодической проверки доступности веб-серверов.

Что делает:

- Читает список URL из YAML-конфигурации
- Проверяет доступность сайтов с заданным интервалом
- Сверяет полученные HTTP-статусы с ожидаемыми
- Измеряет время ответа серверов
- Работает в многопоточном режиме через worker pool
- Поддерживает graceful shutdown

Из чего состоит:

- **main.go** - основной модуль, оркестрация процессов
- **config.go** - загрузка и парсинг конфигурации
- **workerpool/** - пакет для управления воркерами
  - **pool.go** - пул воркеров, управление задачами
  - **worker.go** - обработчик HTTP-запросов
- **config.yaml** - файл конфигурации
- **config_test.go** - тесты конфигурации

Как использовать:

1. Настрой `config.yaml`:

*пример*

```yaml```
interval: 10s
request_timeout: 2s
workers_count: 3
urls:
  - name: "Google"
    url: "https://google.com/"
    expected_status: 200

2. Запусти программу:

bash

go run main.go config.go
*или*
go run .
*или*
./server_checker.exe


Конфигурация:

- `interval` - как часто проверять URL (например: 10s, 1m)

- `request_timeout` - таймаут для HTTP-запросов

- `workers_count` - количество параллельных воркеров

- `urls` - список проверяемых сайтов с ожидаемыми статусами


Пример вывода:


[SUCCESS] - [Google] - Status: 200 (expected status: 200), Time: 150ms
[WRONG STATUS] - [GitHub] - Status: 404 (expected status: 200), Time: 230ms
[ERROR] - [Example] - https://example.com/ - connection refused


Особенности:

- Graceful shutdown (Ctrl+C для остановки)

- Параллельная обработка URL

- Логирование результатов в реальном времени

- Простая конфигурация через YAML
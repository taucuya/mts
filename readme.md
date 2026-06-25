# mws

CLI-утилита для работы с профилями в текущей директории.

## Описание

Программа позволяет:

* создавать профиль;
* получать профиль по имени;
* выводить список профилей;
* удалять профиль;
* выводить справку по доступным командам.

Профили хранятся в YAML-файлах в текущей папке.
Имя профиля соответствует имени файла, а содержимое файла включает два поля:

* `user`
* `project`

## Сборка

Установить зависимости и собрать бинарник:

```bash
go mod tidy
go build -o mws ./cmd/mws
```

После этого в текущей директории появится исполняемый файл `mws`.

## Быстрый старт

```bash
./mws help
./mws profile create --name=test --user=example --project=new-project
./mws profile get --name=test
./mws profile list
./mws profile delete --name=test
```

## Поддерживаемые команды

### 1. Создание профиля

Команда:

```bash
./mws profile create --name=<name> --user=<user> --project=<project>
```

Пример:

```bash
./mws profile create --name=test --user=example --project=new-project
```

Пример вывода:

```text
profile created
```

После выполнения в текущей папке будет создан файл:

```text
test.yaml
```

Содержимое файла:

```yaml
user: example
project: new-project
```

---

### 2. Получение профиля по имени

Команда:

```bash
./mws profile get --name=<name>
```

Пример:

```bash
./mws profile get --name=test
```

Пример вывода:

```text
name: test
user: example
project: new-project
```

---

### 3. Получение списка профилей

Команда:

```bash
./mws profile list
```

Пример:

```bash
./mws profile list
```

Пример вывода:

```text
NAME   USER     PROJECT
test   example  new-project
dev    alice    billing
```

Профили выводятся на основе YAML-файлов, найденных в текущей директории.

---

### 4. Удаление профиля

Команда:

```bash
./mws profile delete --name=<name>
```

Пример:

```bash
./mws profile delete --name=test
```

Пример вывода:

```text
profile deleted
```

После этого файл `test.yaml` будет удален из текущей директории.

---

### 5. Справка

Команда:

```bash
./mws help
```

Пример вывода:

```text
Usage:
  mws profile create --name=<name> --user=<user> --project=<project>
  mws profile get --name=<name>
  mws profile list
  mws profile delete --name=<name>
  mws help

Commands:
  profile create   Create a profile in current directory
  profile get      Show profile by name
  profile list     List profiles in current directory
  profile delete   Delete profile by name
  help             Show this help
```

## Полный пример использования

### Шаг 1. Сборка

```bash
go mod tidy
go build -o mws ./cmd/mws
```

### Шаг 2. Создание профилей

```bash
./mws profile create --name=test --user=example --project=new-project
./mws profile create --name=dev --user=alice --project=billing
```

Вывод:

```text
profile created
profile created
```

### Шаг 3. Просмотр списка

```bash
./mws profile list
```

Вывод:

```text
NAME   USER     PROJECT
dev    alice    billing
test   example  new-project
```

### Шаг 4. Получение конкретного профиля

```bash
./mws profile get --name=dev
```

Вывод:

```text
name: dev
user: alice
project: billing
```

### Шаг 5. Удаление профиля

```bash
./mws profile delete --name=test
```

Вывод:

```text
profile deleted
```

### Шаг 6. Проверка, что профиль удалён

```bash
./mws profile list
```

Вывод:

```text
NAME  USER   PROJECT
dev   alice  billing
```

## Формат хранения

Каждый профиль хранится в отдельном YAML-файле:

* имя профиля `test` -> файл `test.yaml`
* имя профиля `dev` -> файл `dev.yaml`

Пример файла:

```yaml
user: example
project: new-project
```

## Ошибки и граничные случаи

### Профиль уже существует

Команда:

```bash
./mws profile create --name=test --user=example --project=new-project
./mws profile create --name=test --user=another --project=other-project
```

Пример вывода:

```text
profile created
error: profile "test" already exists
```

### Профиль не найден

Команда:

```bash
./mws profile get --name=missing
```

Пример вывода:

```text
error: profile "missing" not found
```

Или:

```bash
./mws profile delete --name=missing
```

Пример вывода:

```text
error: profile "missing" not found
```

### Не переданы обязательные флаги

Команда:

```bash
./mws profile create --name=test --user=example
```

Пример вывода:

```text
flags --name, --user, --project are required
```

Команда:

```bash
./mws profile get
```

Пример вывода:

```text
flag --name is required
```
## Входные данные

CLI принимает следующие входные данные:

* `profile create --name --user --project`
* `profile get --name`
* `profile list`
* `profile delete --name`
* `help`

## Результат работы

Программа:

* создает YAML-файлы профилей в текущей директории;
* читает и выводит данные профиля по имени;
* показывает список профилей;
* удаляет профиль по имени;
* выводит справку по командам.

## Принятые допущения

В реализации использованы следующие допущения:

1. Имя профиля преобразуется в имя файла по шаблону `<name>.yaml`.
2. Все операции выполняются только в текущей директории.
3. Для имени профиля допустимы только:

   * латинские буквы;
   * цифры;
   * символы `-` и `_`.
4. В команде `profile list` читаются файлы с расширениями:

   * `.yaml`
   * `.yml`
5. Если YAML-файл поврежден или не соответствует ожидаемой структуре, команда завершится с ошибкой.
6. Если профиль уже существует, повторное создание не перезаписывает файл, а возвращает ошибку.

## Почему выбрано такое решение

Решение построено на стандартной библиотеке Go и минимальном количестве внешних зависимостей.
Для сериализации и десериализации YAML используется `gopkg.in/yaml.v3`.

## Возможные улучшения

Решение можно расширить:

* добавить команду обновления профиля;
* добавить вывод в JSON;
* добавить рекурсивный поиск профилей;
* игнорировать невалидные YAML-файлы при `profile list`;
* добавить интеграционные тесты CLI.

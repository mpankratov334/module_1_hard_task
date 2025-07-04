# Merge Channel (Fan-In)

## Описание задачи

В этой задаче необходимо реализовать паттерн "Merge Channel" (fan-in) на языке Go.
Идея заключается в том, чтобы объединить несколько входных каналов (типа `<-chan int`) в один выходной канал (`<-chan int`),
в который будут поступать все значения из входных каналов.

Функция должна:

- Принимать любое количество входных каналов.
- Возвращать единый выходной канал, в который будут переданы все элементы из входных каналов.
- После завершения передачи данных из всех входных каналов, закрывать выходной канал.

## Подсказки по реализации


## Как запустить тесты

1. Перейдите в директорию проекта:
   ```bash
   cd <название задачи>
   ```
2. Запустите тесты командой:
   ```bash
   go test -v
   ```

## Тесты проверяют:

- Объединение значений из нескольких входных каналов.
- Корректное закрытие выходного канала после завершения работы всех входных каналов.
- Поведение функции при отсутствии входных каналов.


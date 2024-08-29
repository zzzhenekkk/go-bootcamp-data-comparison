# go-bootcamp-data-comparison
Программа на Go для сравнения данных, закодированных в JSON/XML, и файловых систем. Определяет изменения, такие как добавление, удаление или модификация элементов.

# Go Bootcamp: Day 01 - Сравнение Данных на Go

## Описание

Этот проект представляет собой программу на языке Go для сравнения двух файловых систем или баз данных, закодированных в формате JSON или XML. Программа определяет, какие элементы были добавлены, удалены или изменены между двумя версиями файлов.

## Задачи

1. **Exercise 00: Чтение и анализ данных**
   - Программа должна уметь читать и парсить данные из форматов JSON и XML.
   - Используйте стандартные библиотеки Go для работы с этими форматами.

2. **Exercise 01: Оценка изменений**
   - Программа сравнивает два JSON/XML файла и выводит, какие элементы были добавлены, удалены или изменены.
   - Вывод формируется в виде списка действий: `ADDED`, `REMOVED`, `CHANGED`.

3. **Exercise 02: Послесловие**
   - Программа для сравнения файловых систем, представленных в виде текстовых файлов. Она должна определить, какие файлы были добавлены или удалены между двумя версиями файловой системы.
   - Файлы могут быть большими, поэтому необходимо реализовать эффективное использование памяти.

## Правила

- Программа не должна завершаться с ошибкой на корректных данных.
- Использование внешних зависимостей должно управляться с помощью Go Modules.
- Программа должна компилироваться с помощью команды `go build`.

## Инструкции по сборке

Для сборки программы используйте следующую команду:

```bash
go build
// dup выводит текст каждой строки, которая повторяется в
// стандартном вводе более одного раза, а также кол-во появлений
// конец ввода это Ctrl+D
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)      // создали пустое отображение (словарь, map) строк И целых значений в counts
	input := bufio.NewScanner(os.Stdin) // соберём В БУДУЩЕМ ввод в переменную input через нечто Scanner

	// была идея просто вывести все введенные строки повторно без подсчётов.
	// input_backup := input // копия переменной, чтобы вывести строки
	// вероятно, работа со Scanner как-то искажает его результаты

	fmt.Println("Вот, что ввели:")

	//	for input.Scan() {
	//		fmt.Printf(input.Text())
	//	}

	//	fmt.Printf(input_backup)

	for input.Scan() { // для каждого скана из input. значения начинают поступать из stdin в момент обращения к Scan !
		counts[input.Text()]++ // в словаре нашли строку и у строки прибавили число на единицу
	}
	// игнорируем потенциальные ошибки из input.Err()
	fmt.Println("РЕЗУЛЬТАТЫ")
	for line, n := range counts { // для каждой строки и целого числа из диапазона переменной counts
		if n > 0 { // если целое число (значение по ключу) больше 1, то
			fmt.Printf("%d\t%s\n", n, line) // выводим на экран счётчик S%d, таб, строку %s, перевод строки
		}
	}
	for n, line := range counts {
		fmt.Printf("%d\t%s\n", line, n)
	}
}

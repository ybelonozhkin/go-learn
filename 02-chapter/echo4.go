/* echo4 выводит аргументы командной строки
это из книги. ввёл всё 04.09.2021 */
package main

import (
	"flag"
	"fmt"
	"strings"
)

// создаём переменную-флаг типа пакетом Bool с тремя аргументами
var n = flag.Bool("n", false, "пропуск символа новой строки")

// создаём переменную типа string тоже с тремя аргументами
var sep = flag.String("s", " ", "разделитель")

func main() {
	flag.Parse()                               // парсим аргументы
	fmt.Print(strings.Join(flag.Args(), *sep)) // печатаем аргументы
	if !*n {                                   // если n != True, то печатаем в конце ещё и возврат строки
		fmt.Println()
	}
	println(*n) // мои отладочные выводы
	println(*sep)
}

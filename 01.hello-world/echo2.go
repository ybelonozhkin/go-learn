// echo2 выводит аргументы командной строки (вариант)
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""                  // sep - разделитель, s - строка. := это краткое объявление переменной на основе типа инициализатора
	for i, arg := range os.Args[0:] { // для каждого i и arg из диапазона. Так нельзя перебирать строки
		s += sep + arg // к s добавить разделитель и аргумент
		sep = " "
		fmt.Println(i, os.Args[i])
	}
	// fmt.Println(strings.Join(os.Args[0:], " "))

}

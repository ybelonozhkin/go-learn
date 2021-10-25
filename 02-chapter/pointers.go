// тренировки с указателями 02.09.2021
package main

func main() {
	x := 1  // значение 1
	p := &x // p указывает на адрес х
	c := *p // c указывает на значение по адресу p (да)
	println("x: ", x, &x, *&*&x)
	println("p: ", p)
	println(&x == p)
	println("c: ", c)
	*p = 2 // по указателю присвоили значение
	println("x:", x)
	println("*p =", *p)
}

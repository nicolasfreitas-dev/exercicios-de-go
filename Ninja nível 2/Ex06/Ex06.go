/*
Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
Demonstre estes valores.
*/

package main

import "fmt"

const (
	a = 2023 + iota
	b
	c
	d
	e
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
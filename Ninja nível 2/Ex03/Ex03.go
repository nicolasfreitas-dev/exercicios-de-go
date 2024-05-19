/*
Crie constantes tipadas e não-tipadas.
Demonstre seus valores.
*/

package main

import (
	"fmt"
)

const a int = 10
const b = 20
const c string = "texto"
const d = "20"

func main() {
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
}
/*
Crie um loop utilizando a sintaxe: for {}
Utilize-o para demonstrar os anos desde que você nasceu.
*/

package main

import "fmt"

func main() {
	x := 1999
	y := 2024

	for {
		if (x > y) {
			break
		}

		fmt.Println(x)
		x++
	}
}
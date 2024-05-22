/*
Crie um loop utilizando a sintaxe: for condition {}
Utilize-o para demonstrar os anos desde que você nasceu.
*/

package main

import "fmt"

func main() {
	x := 1999
	y := 2024

	for x <= y {		
		fmt.Println(x)
		x++
	}
}
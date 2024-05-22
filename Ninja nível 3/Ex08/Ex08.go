/*
Crie um programa que utilize a declaração switch, sem switch statement especificado.
*/

package main

import "fmt"

func main() {
	x := 1

	switch {
	case x == 10:
		fmt.Println("no primeiro")
	case x < 1:
		fmt.Println("no segundo")
	case x == 1:
		fmt.Println("no terceiro")
	}
}
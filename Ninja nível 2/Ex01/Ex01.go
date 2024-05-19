/*
Escreva um programa que mostre um número em decimal, binário, e hexadecimal.
*/

package main

import (
	"fmt"
)

func main() {
	x := 10

	fmt.Printf("Decimal: %d \nBinário: %b \nHexadecimal %#x", x, x, x)
}
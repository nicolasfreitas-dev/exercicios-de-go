/*
Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".
*/

package main

import "fmt"

func main() {
	esporteFavorito := "basquete"

	switch esporteFavorito {
	case "karate":
		fmt.Println("esporte de luta")
	case "natação":
		fmt.Println("esporte aquático")
	case "corrida":
		fmt.Println("esporte de resistência")
	case "basquete":
		fmt.Println("seu esporte favorito")
	}
}
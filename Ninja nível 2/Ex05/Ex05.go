/*
Crie uma variável de tipo string utilizando uma raw string literal.
Demonstre-a.
*/

package main

import (
	"fmt"
)

func main() {
	x := `Essa string
			vai ficar
			doidona
			pra testar
			`

	fmt.Printf(x)
}
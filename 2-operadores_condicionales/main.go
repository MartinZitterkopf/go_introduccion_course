package main

import (
	"fmt"
)

func main() {
	edad := 29
	esMayor := 18

	fmt.Println("Operadores")
	fmt.Printf("Su edad es %d años, para entrar debe ser mayor de %d! ", edad, esMayor)
	if edad > esMayor {
		fmt.Printf("Bienvenido!\n")
	} else {
		fmt.Printf("No es bienvenido, regrese cuando tenga %d\n", esMayor)
	}

	if value := funcionPrueba(edad); value == true {
		fmt.Println("Entre al if")
	} else {
		fmt.Println("Sali por el else")
	}

}

// La funcionPrueba recibe un parametro de edad y devuelve un bool
func funcionPrueba(val int) bool {
	var valor bool
	if val > 18 {
		valor = true
	} else {
		valor = false
	}

	switch valor {
	case true:
		fmt.Println("Permitido")
	default:
		fmt.Println("No permitido")
	}

	return valor
}

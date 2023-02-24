package main

import (
	"fmt"
)

func main() {

	// MAPAS
	// instancear un mapa con make
	myMap1 := make(map[int]string)
	myMap1[1] = "uno"
	myMap1[2] = "dos"
	myMap1[3] = "tres"
	myMap1[98] = "valor a borrar"
	myMap1[99] = "valor a cambiar"

	// acceder a una key especifica del mapa para ver su valor
	fmt.Println("imprimir un valor del mapa pasando una clave: ", myMap1[1])

	fmt.Println("imprimir el mapa iniciarl: ", myMap1)

	// borrar una key del mapa
	delete(myMap1, 98)
	fmt.Println("imprimir el mapa luego de borrar: ", myMap1)

	// modificar un valor del mapa
	myMap1[99] = "noventa y nueve"
	fmt.Println("imprimir el mapa luego de modificar valor de clave 99: ", myMap1)

	// retornar valores para saber si la clave existe
	valor, ok := myMap1[66]
	fmt.Println("El valor que tiene es: ", valor, "existe la clave en el mapa? ", ok)

	// instancear mapa con valores por defecto sin utilizar make
	myMap2 := map[string]int{
		"uno":  1,
		"dos":  2,
		"tres": 3,
	}
	fmt.Println("Segundo mapa: ", myMap2)

}

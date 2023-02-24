package main

import (
	"fmt"
)

func main() {
	sum1 := 0
	sum2 := 0
	sum3 := 0
	myArray := []string{"cero", "uno", "dos", "tres", "cuatro", "cinco", "seis", "siete", "ocho", "nueve"}
	myMapa1 := map[int]string{
		0: "cero",
		1: "uno",
		2: "dos",
		3: "tres",
		4: "cuatro",
		5: "cinco",
		6: "seis",
		7: "siete",
		8: "ocho",
		9: "nueve",
	}

	myMapa2 := map[string][]int{
		"cero":   nil,
		"uno":    {1, 2, 3, 4, 5},
		"dos":    {2, 4, 6, 8, 10},
		"tres":   {3, 6, 9, 12, 15},
		"cuatro": {4, 8, 12, 16, 20},
		"cinco":  {5, 10, 15, 20, 25},
	}

	// CICLO FOR
	// Primer metodo, con variable de inicio y fin
	for i := 0; i < 10; i++ {
		sum1++
		// fmt.Println(sum)
	}
	fmt.Println("Primer metodo: ", sum1)

	// Segundo metodo, con valor de salida
	for sum2 < 100 {
		sum2++
	}
	fmt.Println("Segundo metodo: ", sum2)

	// Tercer metodo, con break para salir del bucle
	for {
		if sum3 >= 100 {
			break
		}
		sum3++
	}

	fmt.Println("Tercer metodo: ", sum3)

	// Utilizar ciclo for para recorrer un array, mostrando valor llamando al array por su indice
	fmt.Println("ARRAY: mostrando valor llamando al array por su indice")
	for i := range myArray {
		fmt.Println("Index: ", i, " - Valor: ", myArray[i])
	}

	// Utilizar ciclo for para recorrer un array, mostrando el valor llamando a una segunda variable
	fmt.Println("ARRAY: mostrando valor llamando variable")
	for i, v := range myArray {
		fmt.Println("Index: ", i, " - Valor: ", v)
	}

	// Utilizar ciclo for para recorrer un mapa
	fmt.Println("recorrer un MAPA")
	for key, value := range myMapa1 {
		fmt.Println("Key: ", key, " - Valor: ", value)
	}

	// Recorrer un mapa de array
	fmt.Println("recorrer un ARRAY DE MAPA")
	for key, value := range myMapa2 {
		fmt.Println("key: ", key)
		for _, valueArray := range value {
			fmt.Println("value: ", valueArray)
		}
		fmt.Println()
	}
}

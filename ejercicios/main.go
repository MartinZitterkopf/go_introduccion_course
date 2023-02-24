package main

import (
	"fmt"

	"github.com/MartinZitterkopf/go_introduccion_course/ejercicios/matrix"
)

// EjercicioClase21
var edad int = 29
var carnet bool = true

// --------------------------------

func main() {

	EjercicioClase21(edad, carnet)
	EjercicioClase26()
	EjercicioClase27()
	EjercicioClase28()
	EjercicioClase34()
}

func EjercicioClase21(edad int, lic bool) {
	age := edad
	license := lic

	if license == true && age > 15 {
		fmt.Println("Puede seguir avanzando")
	} else {
		fmt.Println("No puede seguir circulando")
	}
}

func EjercicioClase26() {
	myArray := [5]int{4, 2, 5, 6, 7}

	for i, v := range myArray {
		myArray[i] = v + 20
	}
	// Otra manera de hace rlo mismo
	// for i := range array {
	// 	array[i] += 20
	// }
	fmt.Println("Los valores del array son: ", myArray)
}

func EjercicioClase27() {
	var myArray []string
	fmt.Println("Ingrese numeros por teclado, para terminar ingrese el 0 (cero)")

	for {
		var value string

		fmt.Scanf("%s", &value)
		if value == "0" {
			break
		}

		myArray = append(myArray, value)
	}

	fmt.Println("Los valores del array son: ", myArray)
}

func EjercicioClase28() {
	var myArray []string
	fmt.Println("Ingrese los codigos de los productos, para salir 0 (cero)")

	for {
		var value string
		fmt.Scanln(&value)

		if value == "0" {
			break
		}

		switch value {
		case "10":
			myArray = append(myArray, "notebook")
		case "15":
			myArray = append(myArray, "tv")
		case "21":
			myArray = append(myArray, "heladera")
		case "27":
			myArray = append(myArray, "monitor")
		case "35":
			myArray = append(myArray, "camara")
		default:
			myArray = append(myArray, "notFound")
		}
	}

	fmt.Println("Los productos del array son: ", myArray)
}

func EjercicioClase34() {
	m, err := matrix.New([]float64{2, 7, 8}, []float64{4, 4, 7}, []float64{5, 6, 1})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	m.Print()
}

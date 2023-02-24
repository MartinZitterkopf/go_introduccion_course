package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	// DECLARAR Y LUEGO ASIGNAR VALOR
	var myIntVar int
	myIntVar = -12
	fmt.Printf("Mi variable del tipo %T tiene un valor de: %v \n", myIntVar, myIntVar)
	fmt.Println("La direccion de memoria es: ", &myIntVar)

	var myUintVar uint
	myUintVar = 12
	fmt.Printf("Mi variable del tipo %T tiene un valor de: %v \n", myUintVar, myUintVar)
	fmt.Println("La direccion de memoria es: ", &myUintVar)

	var myStringVar string
	myStringVar = "primer parte del curso Go desde 0"
	fmt.Printf("Mi variable del tipo %T tiene un valor de: %v \n", myStringVar, myStringVar)
	fmt.Println("La direccion de memoria es: ", &myStringVar)

	var myBoolVar bool
	myBoolVar = true
	fmt.Printf("Mi variable del tipo %T tiene un valor de: %v \n", myBoolVar, myBoolVar)
	fmt.Println("La direccion de memoria es: ", &myBoolVar)

	const myIntConst int = 1
	fmt.Printf("Mi variable del tipo %T tiene un valor de: %v \n", myIntConst, myIntConst)

	// ASIGNAR VALOR SIN DECLARAR
	myIntVar2 := 24
	fmt.Printf("Mi variable del tipo %T tiene un valor de: %v \n", myIntVar2, myIntVar2)
	fmt.Println("La direccion de memoria es: ", &myIntVar2)

	// COMPARACION DE MEMORIA BYTES
	var myInt8Var int8
	fmt.Printf("Tipo: %T, bytes: %d, bits: %d \n", myInt8Var, unsafe.Sizeof(myInt8Var), unsafe.Sizeof(myInt8Var)*8)
	var myInt16Var int16
	fmt.Printf("Tipo: %T, bytes: %d, bits: %d \n", myInt16Var, unsafe.Sizeof(myInt16Var), unsafe.Sizeof(myInt16Var)*8)
	var myInt32Var int32
	fmt.Printf("Tipo: %T, bytes: %d, bits: %d \n", myInt32Var, unsafe.Sizeof(myInt32Var), unsafe.Sizeof(myInt32Var)*8)
	var myInt64Var int64
	fmt.Printf("Tipo: %T, bytes: %d, bits: %d \n", myInt64Var, unsafe.Sizeof(myInt64Var), unsafe.Sizeof(myInt64Var)*8)

	// CONVERSION DE TIPO DE DATOS

	{
		// De esta manera abrimos un scope {} en golang

		// Convertimos de un float a un string utilizando el package fmt
		floatVar := 12.08
		fmt.Printf("tipo: %T, valor: %f \n", floatVar, floatVar)
		floatString := fmt.Sprintf("%f", floatVar)
		fmt.Printf("tipo: %T, valor: %v \n", floatString, floatString)

		// Convertimos string utilizando el package strconv
		intVar1, err := strconv.ParseInt("12", 0, 64)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("tipo: %T, valor: %d \n", intVar1, intVar1)

		intVar2, _ := strconv.ParseInt("14", 0, 64)
		fmt.Printf("tipo: %T, valor: %d \n", intVar2, intVar2)
	}

	{
		fmt.Println()
		fmt.Println("Values in ASCII code:")

		// para pasar el valor de codigo ASCII se pasa con ''
		var A byte = 'A'
		var a byte = 'a'

		fmt.Println(A)
		fmt.Println(a)

		// para pasar el codigo ASCII se lo pasamos como numero y lo debemos imprimir como un string
		var R byte = 82
		var s byte = 115

		fmt.Println(string(R))
		fmt.Println(string(s))
	}
}

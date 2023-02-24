package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var myIntVar int
	fmt.Println(myIntVar)
	fmt.Printf("tipo: %T, bytes: %d, bits: %d\n", myIntVar, unsafe.Sizeof(myIntVar), unsafe.Sizeof(myIntVar)*8)

	// ARRAY => TAMAÑO FIJO
	var myArray1 [6]int
	fmt.Println("Por defecto luego de crear un array: ", myArray1)
	myArray1[1] = 1
	myArray1[2] = 2
	myArray1[3] = 3
	fmt.Println("Se asignan valores manualmente", myArray1)

	fmt.Println("Size del array: ", len(myArray1))
	fmt.Printf("tipo: %T, bytes: %d, bits: %d\n", myArray1, unsafe.Sizeof(myArray1), unsafe.Sizeof(myArray1)*8)

	myArray2 := [3]string{"value1", "value2", "value3"}
	fmt.Println("Se asignan valores en la definicion del array", myArray2)

	// SLICE => TAMAÑO DINAMICO
	var mySlice1 []int
	fmt.Printf("Size del slice: %d, value: %v\n", len(mySlice1), mySlice1)

	// Agregar datos al slice
	mySlice1 = append(mySlice1, 10, 20, 30, 40, 50)
	fmt.Printf("Size del slice luego de agregar valores: %d, value: %v\n", len(mySlice1), mySlice1)
	fmt.Printf("tipo: %T, bytes: %d, bits: %d\n", mySlice1, unsafe.Sizeof(mySlice1), unsafe.Sizeof(mySlice1)*8)

	// Tomar una porcion de un array/slice para hacer un slice
	mySlice2 := myArray1[2:4]
	fmt.Println(mySlice2)

	// Si buscamos en memoria se ve que el slice comparte la ubicacion con el array, esto se debe a que el slice crea una vista del vector
	fmt.Println("Array en posicion 2: ", &myArray1[2])                                        // 2
	fmt.Println("Slice en posicion 0, toma el valor del array en posicion 2: ", &mySlice2[0]) // 2

	// instancear slice con make()
	mySlice3 := make([]int, 3)
	fmt.Println(mySlice3)

}

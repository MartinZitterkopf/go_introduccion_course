package main

import (
	"fmt"

	"github.com/MartinZitterkopf/go_introduccion_course/6-funciones/function"
)

func main() {
	fmt.Println(function.AddSuma(15, 42))

	function.RepeatString(3, "Welcome")
	fmt.Println()

	v, err := function.Calculadora(function.MUL, 3, 11)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("value: ", v)
	}

	x, y := function.ValorCuadrado(5)
	fmt.Println("El valor de", x, "al cuadrado es:", y)

	fmt.Println("El valor de la multisuma es: ", function.MultiSumEnteros(2, 5, 9, 8, 7, 8, 7, 8, 9, 6, 3))

	// llamamos a la funcion FactoryOperation y almacenamos en una variable la respuesta que es otra funcion
	fn := function.FactoryOperation(function.SUM)
	// pasamos nuevos parametros a la nueva funcion
	v = fn(2, 4)
	fmt.Println("Suma: ", v)

	// llamamos a la funcion FactoryOperation y almacenamos en una variable la respuesta que es otra funcion
	fn = function.FactoryOperation(function.MUL)
	// pasamos nuevos parametros a la nueva funcion
	v = fn(2, 4)
	fmt.Println("Multiplicacion: ", v)
}

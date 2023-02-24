package function

import (
	"errors"
	"fmt"
)

type Operation int

// iota es como hacer un enum, para utilizarlo todas las contantes deben llamarse en un mismo bloque, de otra manera serian individuales los valores y a todos se les asignaria 0
const (
	SUM Operation = iota
	RES
	DIV
	MUL
)

// Otra forma seria hacerlo manual con constantes independientes
// const SUM Operation = 0
// const RES Operation = 1
// const DIV Operation = 2
// const MUL Operation = 3

// cuando se rretorna un solo valor se puede obviar el parentesis, de hecho lo elimina al guardar el archivo automaticamente
func AddSuma(x int, y int) int {
	return x + y
}

func RepeatString(cant int, value string) {
	for i := 0; i < cant; i++ {
		fmt.Print(value, " ")
	}
}

// cuando se retorna mas de un valor va entre parentesis
func Calculadora(operacion Operation, x float64, y float64) (float64, error) {
	switch operacion {
	case SUM:
		return x + y, nil
	case RES:
		return x - y, nil
	case DIV:
		if y == 0 {
			return 0, errors.New("No se puede dividir por 0")
		}
		return x / y, nil
	case MUL:
		return x * y, nil

	}

	return 0, errors.New("Ah ocurrido un error")
}

// En las funciones tambien se puede declarar las variables de retorno, de esta manera solo se escribe el return y no hay que devolver los valores ahi
func ValorCuadrado(value int) (x int, y int) {
	// como se declararon x e y en la funcion de retorno solo se tiene que agregar el valor
	x = value
	y = value * value
	// como se declararon los valores que se deben retornar, no hace falta especificarlos
	return
}

// values ...int almacenara los valores como si fuera un array
func MultiSumEnteros(values ...int) int {
	var sum int
	for _, value := range values {
		sum += value
	}

	return sum
}

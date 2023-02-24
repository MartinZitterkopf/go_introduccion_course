package main

import (
	"fmt"
	"strconv"
)

func main() {

	// CONVERTIR A STRING
	s := strconv.Itoa(-42) // convertir de entero a string
	fmt.Printf("tipo: %T, valor: %s\n", s, s)

	s = strconv.FormatBool(true) // convertir de booleano a string
	fmt.Printf("tipo: %T, valor: %s\n", s, s)

	s = strconv.FormatFloat(3.1415, 'E', -1, 64) // convertir de flotante a string
	fmt.Printf("tipo: %T, valor: %s\n", s, s)

	s = strconv.FormatInt(-42, 16) // convertir de int a string hexadecimal (16) => para decimal se usa 10
	fmt.Printf("tipo: %T, valor: %s\n", s, s)

	s = strconv.FormatUint(42, 16) // convertir de Uint a string hexadecimal (16) => para decimal se usa 10
	fmt.Printf("tipo: %T, valor: %s\n", s, s)

	//CONVERTIR DESDE STRING
	b, err := strconv.ParseBool("true") // convertir string a booleano
	fmt.Println(b, err)

	f, err := strconv.ParseFloat("3.1415", 64) // convertir string a flotante
	fmt.Println(f, err)

	i, err := strconv.ParseInt("-42", 10, 64) // convertir string a int
	fmt.Println(i, err)

	u, err := strconv.ParseInt("42", 10, 64) // convertir string a Uint
	fmt.Println(u, err)

}

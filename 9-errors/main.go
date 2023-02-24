package main

import (
	"errors"
	"fmt"

	"github.com/MartinZitterkopf/go_introduccion_course/9-errors/operators"
)

func main() {

	var err error
	err = errors.New("this in a error")
	fmt.Println(err)         // println lo transforma automaticamente a string, pero err es tipo error
	fmt.Println(err.Error()) // utilizando Error() lo transformamos a string

	err2 := fmt.Errorf("my format err, string %s, number: %d", "my string", 23)
	fmt.Println(err2)         // println lo transforma automaticamente a string, pero err es tipo error
	fmt.Println(err2.Error()) // utilizando Error() lo transformamos a string

	defer func() {
		fmt.Println("I am defer")
		r := recover()
		if r != nil {
			fmt.Println("No se pudo terminar ejecucion de codigo")
			fmt.Println("Recover in", r)
		}
	}()

	z := operators.Div(10, 2)
	fmt.Println(z)

}

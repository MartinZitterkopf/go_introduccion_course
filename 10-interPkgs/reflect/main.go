package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID        int64
	Email     string
	FirstName string
	LastName  string
	Age       *float64
	Address   Address
}

type Address struct {
	Country string
	State   string
}

func main() {

	var age float64 = 29
	u := User{1, "mzzz@hotmail.com", "Martin", "zzz", &age, Address{"Argentina", "Buenos Aires"}}
	Examine(u)
}

func Examine(data interface{}) {

	if reflect.ValueOf(data).Kind() == reflect.Struct {
		v := reflect.ValueOf(data)
		fmt.Println("Numero de campos: ", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int, reflect.Int32, reflect.Int64:
				myVar := v.Field(i).Int()
				fmt.Printf("Campo: %d, Tipo: %T, Value: %v\n", i, myVar, myVar)
			case reflect.String:
				myVar := v.Field(i).String()
				fmt.Printf("Campo: %d, Tipo: %T, Value: %v\n", i, myVar, myVar)
			case reflect.Pointer:
				fmt.Printf("Campo: %d, Value: %v\n", i, v.Field(i))
			case reflect.Struct:
				if v.Field(i).Type() == reflect.TypeOf(Address{}) {
					myVar := v.Field(i).Interface().(Address)
					fmt.Printf("Campo: %d, Value: %v\n", i, myVar)
					fmt.Println(myVar.Country)
					fmt.Println(myVar.State)
				} else {
					fmt.Println("Dato no soportado: ", v.Field(i).Type())
				}
			default:
				fmt.Println("Dato no soportado: ", v.Field(i).Type())
			}
		}
	} else {
		t := reflect.TypeOf(data)
		v := reflect.ValueOf(data)
		k := t.Kind()
		fmt.Println("Type: ", t)
		fmt.Println("Value: ", v)
		fmt.Println("Kind: ", k)
	}

}

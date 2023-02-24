package main

import "fmt"

type Number interface {
	int | int32 | int64 | float32 | float64
}

func main() {

	v1 := []float64{1.3, 5.55, 3.66, 9.45, 6.78}
	v2 := []int{5, 8, 9, 6, 3, 7, 5, 12, 5}

	fmt.Println(Suma(v1))
	fmt.Println(Suma(v2))

	fmt.Println(SumaConInterface(v1))
	fmt.Println(SumaConInterface(v2))

	anyType(1, 1)
	anyType("a", "a")
	// anyType(1,"a") // error
	anyTypeX2(1, "a")
	comparableType("a", "b")

	customInt := CustomSlice[int]{1, 2, 3, 1, 5, 4, 9, 7}
	fmt.Println(customInt)

	customString := CustomSlice[string]{"hola", "buen dia", "que tal"}
	fmt.Println(customString)

	fmt.Println(MinNumber(1, 7))

	fd := MyGenericStruct[MyFirstData]{StrValue: "Test", Data: MyFirstData{}}
	fd.Data.PrintOne()

	sd := MyGenericStruct[MySecondData]{StrValue: "Test", Data: MySecondData{}}
	//sd.Data.PrintOne()
	sd.Data.PrintTwo()

}

// entre corchetes ponemos al generico, y luego lo llamamos como parametro
func Suma[N int | int32 | int64 | float32 | float64](v []N) N {
	var total N

	for _, tV := range v {
		total += tV
	}
	return total
}

func SumaConInterface[N Number](v []N) N {
	var total N

	for _, tV := range v {
		total += tV
	}
	return total
}

// cuando se utiliza el generico any, lo que hace go es el primer tipo que recibe tiene que ser igual a la sucesion siguiente
func anyType[N any](v1, v2 N) {
	fmt.Println(v1, v2)
	// fmt.Println(v1 == v2) // no permite compraciones el tipo any
}

func anyTypeX2[N1 any, N2 any](v1 N1, v2 N2) {
	fmt.Println(v1, v2)
}

// para comparar valores que no sabemos el tipo se utiliza el tipo de compare
func comparableType[N comparable](v1, v2 N) {
	fmt.Println(v1, v2)
	fmt.Println(v1 == v2)
}

// para hacer un slice custom
type CustomSlice[V int | string] []V

func MinNumber[T Number](x, y T) T {
	if x < y {
		return x
	}
	return y
}

type (
	MyGenericStruct[D any] struct {
		StrValue string
		Data     D
	}

	MyFirstData struct{}

	MySecondData struct{}
)

func (d MyFirstData) PrintOne() {
	fmt.Println("Print first")
}

func (d MySecondData) PrintTwo() {
	fmt.Println("Print second")
}

package main

import "fmt"

type Producto struct {
	ID   int
	Name string
}

func (m Producto) Set(name string) {
	fmt.Printf("address: %p\n", &m)
	m.Name = name
}

func (m *Producto) SetPointer(name string) {
	fmt.Printf("address: %p\n", m)
	m.Name = name
}

func main() {
	var i int
	var iP *int

	i = 34
	iP = &i

	fmt.Println(&i)
	fmt.Println(iP)
	fmt.Println(i)
	fmt.Println(*iP)

	*iP = 1
	fmt.Printf("val i: %d, val pointer i: %d\n", i, *iP)

	myVar := 30
	fmt.Printf("address: %p, values: %v\n", &myVar, myVar)
	increment(myVar)
	fmt.Println(myVar)
	incrementPointer(&myVar)
	fmt.Println(myVar)

	var mySlice []int
	mySlice = append(mySlice, 3, 4, 2)
	fmt.Printf("address: %p, values: %v\n", &mySlice, mySlice)
	fmt.Printf("address0: %p, address1: %p, address2: %p\n", &mySlice[0], &mySlice[1], &mySlice[2])

	updateSlice(mySlice)
	println(mySlice)

	myStruct := &Producto{ID: 1, Name: "test"}
	fmt.Printf("addrs: %p\n", myStruct)
	fmt.Printf("id: %d, name: %s\n", myStruct.ID, myStruct.Name)
	updateStruct(myStruct)
	fmt.Printf("id: %d, name: %s\n", myStruct.ID, myStruct.Name)

	myStruct.Set("test 3")
	fmt.Printf("id: %d, name: %s\n", myStruct.ID, myStruct.Name)
	myStruct.SetPointer("test 4")
	fmt.Printf("id: %d, name: %s\n", myStruct.ID, myStruct.Name)

}

func increment(val int) {
	fmt.Println(&val)
	val++
}

func incrementPointer(val *int) {
	fmt.Println(val)
	*val++
}

func updateSlice(val []int) {
	fmt.Printf("address: %p\n", val)
	fmt.Printf("address 0: %p, address 0: %p,address 0: %p\n", &val[0], &val[1], &val[2])
	val[0] = 12
}

func updateStruct(val *Producto) {
	fmt.Printf("address in function: %p\n", val)
	val.ID = 2
	val.Name = "test2"
}

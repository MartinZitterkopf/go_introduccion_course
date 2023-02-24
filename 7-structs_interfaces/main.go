package main

import (
	"encoding/json"
	"fmt"

	"github.com/MartinZitterkopf/go_introduccion_course/7-structs_interfaces/struct_interface/structs"
	"github.com/MartinZitterkopf/go_introduccion_course/7-structs_interfaces/struct_interface/vehicles"
)

func main() {
	// COMIENZA STRUCT
	// primer forma de utilizar estructuras con palabra reservada var
	var prod1 structs.Product
	prod1.ID = 1
	prod1.Name = "test1"
	prod1.Type = structs.Type{
		ID:          1,
		Code:        "B+",
		Description: "Electrodomestico",
	}
	prod1.Price = 100

	// segunda forma de utilizar estructuras con declaracion y asignacion de variables
	prod2 := structs.Product{}
	prod2.ID = 2
	prod2.Name = "test2"
	prod2.Type = structs.Type{
		ID:          2,
		Code:        "B+",
		Description: "Electrodomestico",
	}
	prod2.Price = 50
	prod2.Tags = []string{"Linea blanca", "hogar", "eco-friendly"}

	// prod3 := structs.Product{3, "test3", structs.Type{}, 0, 75, nil} // no es una buena practica por eso resalta en amarillo

	prod4 := structs.Product{
		ID:   4,
		Name: "test4",
		Type: structs.Type{
			ID:          3,
			Code:        "B+",
			Description: "Electrodomestico",
		},
		Price: 90,
		Count: 2,
		Tags:  []string{"Linea blanca", "hogar", "eco-friendly"},
	}

	prod5 := structs.Product{
		ID:   5,
		Name: "test5",
		Type: structs.Type{
			ID:          5,
			Code:        "B+",
			Description: "Electrodomestico",
		},
		Price: 190,
		Count: 3,
		Tags:  []string{"Linea blanca", "hogar", "eco-friendly"},
	}
	prod6 := structs.Product{
		ID:   6,
		Name: "test4",
		Type: structs.Type{
			ID:          6,
			Code:        "B+",
			Description: "Electrodomestico",
		},
		Price: 230,
		Count: 1,
		Tags:  []string{"Linea blanca", "hogar", "eco-friendly"},
	}
	fmt.Println(prod1)
	fmt.Println(prod2)
	// fmt.Println(prod3)

	// convertir la estructura en un json en formato byte
	v, err := json.Marshal(prod4)
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
	} else {
		// FORMATO BYTE
		// fmt.Println(v)
		// string(v) convierte el formato de byte a string
		fmt.Println(string(v))
		fmt.Println("Precio total: ", prod4.TotalPrice())
	}

	// el uso de setter en GO es mala practica
	prod4.SetName("no realizar setters")
	prod4.AddTags("tag1", "tag2")
	fmt.Println(prod4)

	car := structs.NewCar(121212)
	car.AddProducts(prod4, prod5, prod6)

	fmt.Println("PRODUCTOS DEL CARRITO")
	fmt.Println("TOTAL DE PRODUCTOS: ", len(car.Products))
	fmt.Printf("Total Car: $%.2f\n", car.Total())

	// COMIENZA INTERFACE
	fmt.Println()
	fmt.Println("VEHICLES")

	carV := vehicles.Car{Time: 120}
	fmt.Println(carV.Distance(), "km/h")

	vArray := []string{"CAR", "MOTO", "TRUCK"}
	dist := 400

	var d float64
	for _, v := range vArray {
		fmt.Printf("vehicle: %s\n", v)
		vehiculo, err := vehicles.NewVehicle(v, dist)
		if err != nil {
			fmt.Println("Error: ", err)
			fmt.Println()
			continue // continue lo que hace es terminar la ejecucion que se esta haciendo y seguir con la siguiente del for
		}

		distance := vehiculo.Distance()
		fmt.Printf("Distance: %.2f\n", distance)
		fmt.Println()
		d += distance
	}

	fmt.Printf("Total distance: %.2f\n", d)
	fmt.Println()
}

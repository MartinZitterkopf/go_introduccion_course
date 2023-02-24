package vehicles

import "fmt"

type Vehicle interface {
	Distance() float64
}

type Car struct {
	Time int
}

type Moto struct {
	Time int
}

type Truck struct {
	Time int
}

const (
	CarVehicle  = "CAR"
	MotoVehicle = "MOTO"
	TruckV      = "TRUCK"
)

func NewVehicle(v string, time int) (Vehicle, error) {
	switch v {
	case CarVehicle:
		return &Car{Time: time}, nil
	case MotoVehicle:
		return &Moto{Time: time}, nil
	case TruckV:
		return &Truck{Time: time}, nil
	}
	return nil, fmt.Errorf("no existe el vehiculo")
}

func (car *Car) Distance() float64 {
	// para el ejemplo se presupone que la velocidad es 100 km/h por tal motivo se dividen los minutos del tiempo sobre 60
	return 100 * float64(car.Time/60)
}

func (mot *Moto) Distance() float64 {
	// para el ejemplo se presupone que la velocidad es 120 km/h por tal motivo se dividen los minutos del tiempo sobre 60
	return 120 * float64(mot.Time/60)
}

func (tru *Truck) Distance() float64 {
	// para el ejemplo se presupone que la velocidad es 120 km/h por tal motivo se dividen los minutos del tiempo sobre 60
	return 70 * float64(tru.Time/60)
}

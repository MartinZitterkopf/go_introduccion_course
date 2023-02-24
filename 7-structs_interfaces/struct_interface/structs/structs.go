package structs

type Product struct {
	// `json:"string"` lo que hace es convertir el ID a formato json como "id"
	ID    uint     `json:"id"`
	Name  string   `json:"name"`
	Type  Type     `json:"type"`
	Count uint16   `json:"count"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

type Type struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type Car struct {
	ID       uint      `json:"id"`
	UserID   uint      `json:"user_id"`
	Products []Product `json:"products"`
}

// METODOS
// para definir metodos, se debe agregar una variable que va a representar la estructura entre la palabra func y el nombre de la funcion
func (prod Product) TotalPrice() float64 {
	return float64(prod.Count) * prod.Price
}

// SET O GET SON MALA PRACTICA NO REALIZAR
func (prod *Product) SetName(name string) {
	prod.Name = name
}

func (prod *Product) AddTags(tags ...string) {
	prod.Tags = append(prod.Tags, tags...)
}

func (car *Car) AddProducts(products ...Product) {
	car.Products = append(car.Products, products...)
}

func (car Car) Total() float64 {
	var total float64
	for _, car := range car.Products {
		total += car.TotalPrice()
	}
	return total
}

func NewCar(userID uint) Car {
	return Car{UserID: userID}
}

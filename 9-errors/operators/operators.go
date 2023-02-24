package operators

import "fmt"

func Div(x, y float64) float64 {
	defer func() {
		fmt.Println("I am defer Div function")
	}()
	return x / y
}

package function

// funcion que devuelva otra funcion
func FactoryOperation(operacion Operation) func(float64, float64) float64 {
	switch operacion {
	case SUM:
		return sum
	case RES:
		return res
	case DIV:
		return div
	case MUL:
		return mul
	}

	return nil
}

func sum(a float64, b float64) float64 {
	return a + b
}

func res(a float64, b float64) float64 {
	return a - b
}

func mul(a float64, b float64) float64 {
	return a * b
}

func div(a float64, b float64) float64 {
	if b == 0 {
		return 0
	}

	return a / b
}

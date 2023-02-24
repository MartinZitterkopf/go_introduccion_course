package main

import "github.com/MartinZitterkopf/example_external/matrix"

func main() {

	input := matrix.NewM([]float64{2, 2, 2},
		[]float64{3, 3, 4},
		[]float64{4, 4, 2},
		[]float64{2, 5, 6},
		[]float64{1, 7, 3})

	input.Print()

}

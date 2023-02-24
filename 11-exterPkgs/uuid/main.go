package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	// New().String() da el mismo resultado que NewString()
	id1 := uuid.New().String()
	fmt.Println("id1: ", id1)

	id2 := uuid.NewString()
	fmt.Println("id2: ", id2)

	uid := uuid.New()
	fmt.Println("uid.Version(): ", uid.Version())
	fmt.Println("uid.String(): ", uid.String())

}

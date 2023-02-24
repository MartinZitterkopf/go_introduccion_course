package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	log.Println("Es parecido al Println de 'fmt' pero agrega marca temporal al msj")
	// fatal interrumpe la ejecucion
	log.Fatal("Es parecido al Print de 'fmt' pero agrega marca temporal al msj y luego llama al exit e interrumpe la ejecucion")
	// panic interrumpe la ejecucion y llama a panic
	log.Panic("Es parecido al Print de 'fmt' pero agrega marca temporal al msj y luego ejecuta un panic con la descripcion de la ruta donde esta el error")

	// GENERAR UN ARCHIVO LOG para guardar los eventos

	// Generamos una fecha para almacenar el nombre del archivo
	date := time.Now()

	file, err := os.Create(fmt.Sprintf("%d.log", date.UnixNano())) // UnixNano() convierte la fecha a nanosegundos para que sea un numero unico
	if err != nil {
		panic(err)
	}

	l := log.New(file, "success: ", log.LstdFlags|log.Lshortfile)
	l.Println("test 1")
	l.Println("test 2")
	l.Println("test 3")
	l.Println("test 4")

	l.SetPrefix("error: ")
	l.Println("test 1")
	l.Println("test 2")
	l.Println("test 3")
	l.Println("test 4")
}

package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
		// el package os proporciona exit para que se interrumpa el codigo de lo contrario si hubo error seguiria igual la ejecucion
		os.Exit(1) // termina el programa si termina correctamente se pasa(0), sino se utiliza otro valor del (1 al 125)
	}

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		fmt.Println(err)
		// el package os proporciona exit para que se interrumpa el codigo de lo contrario si hubo error seguiria igual la ejecucion
		os.Exit(1) // termina el programa si termina correctamente se pasa(0), sino se utiliza otro valor del (1 al 125)
	}

	fmt.Printf("read %d bytes: %q\n", count, data[:count])

	// MANEJO DE VARIABLES DE ENTORNO
	env := os.Getenv("USERNAME")
	fmt.Println("ver una variable de entorno USERNAME: ", env)

	// os.Setenv("ENV_EXISTS", "") // generar o modificar variable de entorno
	// env, ok := os.LookupEnv("ENV_EXISTS")
	// fmt.Println(env, ok) // devuelve la variable de entorno y si existe o no

	// GENERAR CONFIGURACION A BASE DE DATOS
	// os.Setenv("DB_USERNAME", "martinz")
	// os.Setenv("DB_PASSWORD", "123456789")
	// os.Setenv("DB_HOST", "127.0.0.1")
	// os.Setenv("DB_PORT", "3006")
	// os.Setenv("DB_NAME", "root")

	// dbURL := os.ExpandEnv("mongodb://${DB_USERNAME}:${DB_PASSWORD}@$DB_HOST:$DB_PORT/$DB_NAME") // ejemplo no existe
	// fmt.Println(dbURL)

	// MANEJAR DIRECTORIOS
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		// el package os proporciona exit para que se interrumpa el codigo de lo contrario si hubo error seguiria igual la ejecucion
		os.Exit(1) // termina el programa si termina correctamente se pasa(0), sino se utiliza otro valor del (1 al 125)
	}
	fmt.Println("estoy en el directorio: ", dir)

}

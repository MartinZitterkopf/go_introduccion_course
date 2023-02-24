package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// sino utilizamos el paquete de godotenv, no se muestra nada porque no lo reconoce al archivo .env
	fmt.Println(os.Getenv("MY_ENV1"))
	fmt.Println(os.Getenv("MY_ENV2"))
	fmt.Println(os.Getenv("MY_ENV3"))

	// al ejecutar el paquete ahora si muestra los valores que se encontraban dentro del archivo .env (que es el que busca por defecto donde se encuentra el main.go)
	// metodo 1 (en caso de tener varios godotenv.Load se carga siempre el archivo del primero)
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(os.Getenv("MY_ENV1"))
	fmt.Println(os.Getenv("MY_ENV2"))
	fmt.Println(os.Getenv("MY_ENV3"))

	// si queremos cambiar la carpeta donde se encuentra el enviroment tenemos que pasarlo como parametro
	// metodo 2 (en caso de tener varios godotenv.Load se carga siempre el archivo del primero)
	if err := godotenv.Load("nuevo/.var"); err != nil {
		fmt.Println(err)
	}

	fmt.Println(os.Getenv("MY_ENV1"))
	fmt.Println(os.Getenv("MY_ENV2"))
	fmt.Println(os.Getenv("MY_ENV3"))

	// leer el archivo de enviroment
	myEnv, err := godotenv.Read("nuevo/.var")
	fmt.Println(err)
	fmt.Println(myEnv)

	// Si se quiere cargar por algun motivo otra archivo se utiliza godotenv.overload
	if err := godotenv.Overload("nuevo/.var"); err != nil {
		fmt.Println(err)
	}

	fmt.Println(os.Getenv("MY_ENV1"))
	fmt.Println(os.Getenv("MY_ENV2"))
	fmt.Println(os.Getenv("MY_ENV3"))

}

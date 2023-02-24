package main

import (
	"fmt"
	"time"
)

/**
canal <- valor // se envia el valor al canal, ejemplo: se utilizo con variable c dentro del canal myProcessWithChannel()
variable := canal // se envia desde el canal hacia la variable, ejemplo: se utilizo con variable result
**/

func main() {

	// para utilizar el goroutimes ponemos adelante de la funcion la palabra clave gos
	go myProcess("A")
	go myProcess("B")

	time.Sleep(3 * time.Second) // Se agrega para darte tiempo a los procesos para correr antes que finalice el main

	myFirstChannel := make(chan string) // la palabra clave chan es para definir un canal

	go myProcessWithChannel("C", myFirstChannel) // aca llamamos al canal pero no lo utilizamos

	result := <-myFirstChannel // esta variable recupera lo emitido por el canal, cada vez que se usa se hace con la flecha <- que indica que va del canal a la variable
	fmt.Println(result)
	close(myFirstChannel) // cerramos el canal

	channelA := make(chan string)
	channelB := make(chan string)

	go myProcessWithChannel("D", channelA)
	go myOtherProcessWithChannel("E", channelB)

	resultA := <-channelA
	fmt.Println(resultA)

	resultB := <-channelB
	fmt.Println(resultB)

	close(channelA)
	close(channelB)
}

func myProcessWithChannel(p string, c chan string) {
	for i := 0; i < 20; i++ {
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("process: %s - num: %d\n", p, i)
	}
	c <- "ok" // enviamos "ok" cuando termina todo para avisar que el canal finalizo
}

func myOtherProcessWithChannel(p string, c chan string) {
	for i := 0; i < 10; i++ {
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("process: %s - num: %d\n", p, i)
	}

	c <- "ok" // enviamos "ok" cuando termina todo para avisar que el canal finalizo
}

func myProcess(p string) {
	for i := 0; i < 15; i++ {
		time.Sleep(150 * time.Millisecond) // se agrego para darle un cierto deley entre cada iteracion
		fmt.Printf("process: %s - num: %d\n", p, i)
	}
}

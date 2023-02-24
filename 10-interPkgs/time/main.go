package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now() // fecha actual al momento de llamada
	p("fecha actual: ", now)

	then := time.Date(1993, 8, 12, 3, 0, 0, 0, time.UTC) // setear fecha a eleccion
	p("fecha seteada manualmente: ", then)

	then1 := then.Add(time.Hour * 1) // agregar tiempo a uma fecha dada, ya sea año, mes, dia, hora, minutos, segundos
	p("se agrego una hora a la fecha seteada manualmente: ", then1)

	then2 := then.Add(-time.Hour * 1) // restar tiempo a uma fecha dada, ya sea año, mes, dia, hora, minutos, segundos
	p("se resto una hora a la fecha seteada manualmente: ", then2)

	// datos de fecha parciales
	p("parcial del año: ", then.Year())
	p("parcial del mes: ", then.Month())
	p("parcial del dia: ", then.Day())
	p("parcial de la hora: ", then.Hour())
	p("parcial de los minutos: ", then.Minute())
	p("parcial de los segundos: ", then.Second())
	p("parcial de los nanosegundos: ", then.Nanosecond())
	p("parcial de la localizacion: ", then.Location())
	p("parcial del dia la semana: ", then.Weekday())

	// comparar fechas
	p("comparamos la fecha then(seteada) y la fecha now(actual): ")
	p("then(seteada) es antes que now(actual)", then.Before(now))
	p("then(seteada) es despues que now(actual)", then.After(now))
	p("then(seteada) es igual que now(actual)", then.Equal(now))

	// calcular la diferencia entre las fechas
	diff := now.Sub(then)
	p("diferencia entre las fechas now(actual) y then(seteada)", diff)
	p("diferencia entre las fechas now(actual) y then(seteada) en horas", diff.Hours())
	p("diferencia entre las fechas now(actual) y then(seteada) en minutos", diff.Minutes())
	p("diferencia entre las fechas now(actual) y then(seteada) en segundos", diff.Seconds())

}

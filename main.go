package main

import (
	"log"
	"twittor-rodrigo/bd"
	"twittor-rodrigo/handlers"
)

func main() {
	if bd.ChequeoConnection() == false {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}

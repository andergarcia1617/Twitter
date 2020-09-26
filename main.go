package main

import (
	"log"

	"github.com/andergarcia1617/Twitter/bd"
	"github.com/andergarcia1617/Twitter/handles"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handles.Manejadores()
}

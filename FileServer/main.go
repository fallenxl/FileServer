package main

import (
	"fmt"
	"net"
)

func main() {

	//inicializamos el servidor
	s := newServer()
	go s.init()

	ln, err := net.Listen("tcp", ":8282")
	if err != nil {
		return
	}

	defer ln.Close()
	fmt.Println("Servidor escuchando en el puerto 8282")
	fmt.Println("<----------------------Log---------------------->")

	// bucle infinito para que siempre este escuchando peticiones
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		//se crea el usuario nuevo
		//se crea con goroutine para no detener el bucle y siga aceptando peticiones
		go s.newUser(conn)
	}

}

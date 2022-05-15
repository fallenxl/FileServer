package main

import (
	"fmt"
	"net"
	"strings"
)

type server struct {
	channel  map[string]*channel
	commands chan command
}

//se crea un nuevo servidor
func newServer() *server {
	return &server{
		channel:  make(map[string]*channel),
		commands: make(chan command),
	}
}

//se ejecuta cuando se conecta un nuevo usuario
func (s *server) newUser(conn net.Conn) {
	u := &user{
		conn:    conn,
		name:    "Guest",
		command: s.commands,
	}
	u.menu(s)
	u.input()

}

//Acciones del usuario
func (s *server) init() {
	for cmd := range s.commands {
		switch cmd.id {
		case cmd_user:
			s.user(cmd.user, cmd.arg)
		case cmd_list:
			s.list(cmd.user)
		case cmd_suscribe:
			s.suscribe(cmd.user, cmd.arg)
		case cmd_send:
			s.sendFile(cmd.user, cmd.arg)
		case cmd_close:
			s.close(cmd.user)
		}
	}
}

//Cambiar el nombre de usuario
func (s *server) user(u *user, args []string) {
	if len(args) < 2 {
		u.warn("Registre un nombre de usuario")
		return
	}

	u.name = args[1]
	u.msg(fmt.Sprintf("Username: %s", u.name))
}

//suscribirse a un canal
func (s *server) suscribe(u *user, args []string) {

	channelName := args[1]
	r, ok := s.channel[channelName]

	//si el canal no existe, se crea
	if !ok {
		r = &channel{
			name:    channelName,
			members: make(map[net.Addr]*user),
		}
		s.channel[channelName] = r
	}

	r.members[u.conn.RemoteAddr()] = u
	s.quit(u)     //se desconecta del canal actual
	u.channel = r //y se conecta al nuevo

	u.msg(fmt.Sprintf("Has entrado al canal %s", r.name))
	u.channel.broadcast(u, u.name+" ha entrado al canal")
}

//salir de un canal
func (s *server) quit(u *user) {
	if u.channel != nil {
		u.left(" ha salido del canal")
		delete(u.channel.members, u.conn.RemoteAddr())
	}

}

//mostrar lista de canales disponibles
func (s *server) list(u *user) {
	var cnl []string
	// lee el mapa de canales
	for name := range s.channel {
		//los guarda en el slice cnl
		cnl = append(cnl, name)
	}

	//imprime los canales disponibles
	u.warn(fmt.Sprintf("Canales: %s", strings.Join(cnl, ", ")))
}

//salir del servidor
func (s *server) close(u *user) {
	u.left(" ha salido del servidor")
	s.servMsg(fmt.Sprintf("Se ha desconectado %s", u.conn.RemoteAddr().String()))
	u.conn.Close() //cierra la conexion cliente / servidor
}

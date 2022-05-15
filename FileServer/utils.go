package main

import (
	"fmt"
	"time"
)

//if error
func (u *user) ifError(err error) {
	if err != nil {
		u.error(err)
	}
}

//msg-server
func (s *server) servMsg(msg string) {
	time := fmt.Sprintf("%v:%v:%v", time.Now().Hour(), time.Now().Minute(), time.Now().Second())
	fmt.Printf("%s -- %s \n", time, msg)
}

//msg-user
func (u *user) msg(msg string) {
	u.conn.Write([]byte("-- " + msg + "\n"))
}

//warning-user
func (u *user) warn(msg string) {
	u.conn.Write([]byte("!--" + msg + "\n"))
}

//error-user
func (u *user) error(err error) {
	u.conn.Write([]byte("x-- error:" + err.Error() + "\n"))
}

// mensaje al salir
func (u *user) left(msg string) {
	u.warn(msg)
	u.channel.broadcast(u, u.name+msg)
}

// menu de inicio
func (u *user) menu(s *server) {
	u.msg("Lista de comandos: >user, >suscribe, >list, >send, >close")
	u.msg("Registra tu nombre de usuario (command: >user)")
	s.servMsg(fmt.Sprintf("Se ha conectado %s", u.conn.RemoteAddr().String()))
}

package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type user struct {
	conn    net.Conn
	name    string
	channel *channel
	command chan<- command
}

func (u *user) input() {
	for {
		//captura el mensaje desde consola hasta que se presione Enter
		msg, err := bufio.NewReader(u.conn).ReadString('\n')
		if err != nil {
			return
		}

		//elimina salto de lineas
		msg = strings.Trim(msg, "\r\n")

		//divide la entrada comando / argumento
		args := strings.Split(msg, " ")

		//capturamos el comando de la entrada
		cmd := strings.TrimSpace(args[0])

		switch cmd {
		case ">user":
			u.command <- command{ //envia datos de tipo command al canal command (user)
				id:   cmd_user,
				user: u,
				arg:  args,
			}
		case ">list":
			u.command <- command{
				id:   cmd_list,
				user: u,
			}
		case ">suscribe":
			u.command <- command{
				id:   cmd_suscribe,
				user: u,
				arg:  args,
			}
		case ">send":
			u.command <- command{
				id:   cmd_send,
				user: u,
				arg:  args,
			}
		case ">close":
			u.command <- command{
				id:   cmd_close,
				user: u,
			}
		default:
			u.error(fmt.Errorf("comando '%v' desconocido", cmd))
		}
	}
}

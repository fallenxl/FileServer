package main

type commandID int

const (
	cmd_user     commandID = iota // funciona como iterador = 0
	cmd_list                      //1
	cmd_suscribe                  //2
	cmd_members                   //3
	cmd_send                      //4
	cmd_close                     //5
)

type command struct {
	id   commandID
	user *user
	arg  []string
}

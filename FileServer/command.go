package main

type commandID int

const (
	cmd_user     commandID = iota + 1 // funciona como iterador = 1
	cmd_list                          //2
	cmd_suscribe                  	  //3
	cmd_members                       //4
	cmd_send                          //5
	cmd_close                         //6
)

type command struct {
	id   commandID
	user *user
	arg  []string
}

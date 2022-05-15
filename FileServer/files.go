package main

import (
	"fmt"
	"io"
	"os"
)

//directorios
var dirSend string = "./files/"           //directorio de donde se envia el archivo (usuario)
var dirReceive string = "./FilesClients/" //directorio donde se descargara el archivo (usuario)
var dirServer string = "./FilesServer/"   //directorio de donde se obtendra (servidor)

func (s *server) sendFile(u *user, args []string) {

	// verificar si el usuario esta en un canal
	if u.channel == nil {
		u.warn("Entra a un canal o crea uno con el comando '>suscribe <name>'")
		return
	}

	// nombre del archivo
	fileName := args[1]

	// verifica si el directorio existe
	_, err := os.Stat(dirServer + u.channel.name)
	if os.IsNotExist(err) {
		// si no existe lo crea
		os.MkdirAll(dirServer+u.channel.name+"/", 0777)
	}

	// abre el archivo que ha enviado el remitente
	file, err := os.Open(dirSend + fileName)
	if err != nil {
		u.warn(fmt.Sprintf("El archivo '%s' no existe en el directorio", fileName))
		return
	}

	defer file.Close()

	// se crea el archivo en el directorio del servidor, si no existe lo crea
	rc, err := os.OpenFile(dirServer+u.channel.name+"/"+fileName, os.O_RDWR|os.O_CREATE, 0666)
	u.ifError(err)

	defer rc.Close()
	u.ifError(err)

	// se copia el archivo
	_, err = io.Copy(rc, file)
	u.ifError(err)

	//funcion para enviar informacion a los miembros arch.(channel.go)
	u.channel.broadcastFile(u, fileName)
	u.msg("Has enviado un archivo")
}
func (u *user) receiveFile(fileName string) {

	//verifica si el directorio existe
	_, err := os.Stat(dirReceive + u.name)
	if os.IsNotExist(err) {
		//si no existe, se crea
		os.MkdirAll(dirReceive+u.name+"/", 0777)
	}

	//verificamos que el archivo exista en el servidor
	file, err := os.Open(dirServer + u.channel.name + "/" + fileName)
	if err != nil {
		u.warn(fmt.Sprintf("El archivo '%s' no existe en el directorio", fileName))
		return
	}

	defer file.Close()

	//se crea el archivo en el directorio
	rc, err := os.OpenFile(dirReceive+u.name+"/"+fileName, os.O_RDWR|os.O_CREATE, 0666)
	u.ifError(err)

	defer rc.Close()

	//se copia la informacion
	_, err = io.Copy(rc, file)
	u.ifError(err)

}

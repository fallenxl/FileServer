=======================================================================================
                             File Server basado en TCP
=======================================================================================

<----------------------------- Ejecutar Proyecto ------------------------------------->
- VS Code: cree una nueva terminal y dentro del directorio ejecute los comandos go build . 
luego ./fileserver, o el nombre de la carpeta que le haya asignado, si lo requiere ejecute
go mod init para que funcione correctamente.

- CMD (Windows): acceda al directorio donde esta el archivo, y ejecute go build .
luego fileserver, o el nombre de la carpeta que le haya asignado.

<---------------------------------- Comandos ----------------------------------------->

=======================================================================================
|          Comando       |                         Funcion                            |
=======================================================================================
|     >user <name>       |                cambiar de nombre de usuario.               |
|     >suscribe <chan>   | se puede utilizar para unirse a un canal o crear uno nuevo.|
|     >list              |          muestra la lista de canales disponibles.          |
|     >send <file name>  |        enviar el archivo dentro del directorio /files      |
|     >close             |                  desconectarse del servidor                |
=======================================================================================

<------------------------------------ Notas ------------------------------------------>

- Si quiere cambiar los directorios, ajustelo a su preferencia (files.go line: 10)
 

Realizado por: Axl Enrique Santos Hernandez   Mayo 2022
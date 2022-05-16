=======================================================================================
                             File Server basado en TCP
=======================================================================================

<----------------------------- Ejecutar Proyecto ------------------------------------->
- VS Code: cree una nueva terminal y dentro del directorio ejecute los comandos go build . 
luego ./fileserver, o el nombre de la carpeta que le haya asignado, si lo requiere ejecute
go mod init para que funcione correctamente.

- CMD (Windows): acceda al directorio donde esta el archivo, y ejecute go build .
luego fileserver, o el nombre de la carpeta que le haya asignado.

- Usuario: puede acceder desde la consola con el comando <telnet localhost 8282> o puede ajustar 
el puerto a su preferencia (main.go line: 14)

- Enviar archivos: se ha creado una carpeta files dentro del proyecto por el cual se enviaran los archivos
a los usuarios conectados en el canal, se utilizo la carpeta files con intenciones de pruebas para facilitar
la demostracion, pero se puede modificar la direccion del directorio a su preferencia.

Para enviar archivos debe estar dentro de un canal, se sugiere que registre un nombre de usuario
para evitar sobre-escritura de archivos en las carpetas de los usuarios.

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

package main

import (
	TDAHash "main/Hash"
	"main/lectura"
	user "main/usuarios"

	"bufio"
	"os"
	"strings"
)

var usuarios = os.Args[1]

var Posts = TDAHash.CrearHash[int, user.Post]()       // [clave = ID: valor = POST]
var Users = TDAHash.CrearHash[string, user.Usuario]() // [clave = nombre: valor = Usuario]

func main() {
	lista_Usuarios := lectura.LeerUsuarios(usuarios)

	for index, nombre := range lista_Usuarios {
		Users.Guardar(nombre, user.CrearUsuario(nombre, index))
	}

	var UsuarioActivo string
	var ContadorPost int
	S := bufio.NewScanner(os.Stdin)

	for S.Scan() {
		lectura.LecturaDeComandos(strings.Fields(S.Text()), Users, Posts, &UsuarioActivo, &ContadorPost)
	}
}

/*
 ⠄⠄⠄⠄⠄⠄⠄⢀⣟⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣶⢔⢙⣷⡄⠄⠄⠄⠈⣿⣿⣿
 ⠄⠄⠄⠄⠄⢠⡎⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⣿⣿⣧⡄⠄⠄⠄⢸⣇⡚
 ⠄⠄⠄⠄⢀⣣⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡆⠄⠄⢸⣧⣽
 ⠄⠄⠄⠚⠛⠿⠿⠿⣿⣿⣿⠿⠛⠛⠛⠿⣿⣿⣿⣿⣿⣿⣿⠿⠿⠿⢿⣧⠄⢠⠾⠿⢿
 ⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠉⠄⠄⠄⠄⠄⠉⠉⠛⠛⠉⠉⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢸
 ⠄⠄⢸⣇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣠⣤⡄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢰⣶⣾
 ⣿⣄⣿⣿⣧⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣴⣿⣿⣿⠄⠄⠄⠄⠄⠄⠄⠄⢀⣼⣿⣿
 ⣿⣿⣿⣿⣿⣿⣷⣶⣤⣤⣀⣀⣀⣀⣀⣠⣾⣿⣿⣿⣿⣧⠄⠄⠄⠄⢀⣀⣰⠲⠶⠶⠰
 ⡿⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣶⣿⣿⣿⣿⣿⡇⡐⠲⠴⣦
 ⠄⠈⣿⣿⣿⣿⣿⣿⣿⣿⣟⣿⣿⣿⣿⣯⡛⠛⠿⡿⠿⠛⠋⠄⠙⠋⠉⠁⣤⢸⣅⣤⢸
 ⠄⠄⠸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣤⡀⠄⢀⣀⣀⠄⠄⠄⠄⠈⠁⠉⠄⠄⠈
 ⠄⠄⢰⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣛⣛⠛⠛⠛⠉⣁⣀⡔⠄⠄⠄⠄⠄⠄⢰⣶
 ⠄⢀⣾⣿⣿⡿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠿⠛⢉⣿⣿⠟⠄⠈⠄⠄⠄⠄⠄⠘⠛
 ⣴⣿⣿⣿⣿⣿⣎⢻⣿⣿⣿⣿⣿⣿⣿⣿⣧⣀⣠⣤⣾⣿⡿⠄⠄⠄MESSIRVE
*/

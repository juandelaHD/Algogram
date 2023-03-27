package lectura

import (
	"fmt"
	TDAHash "main/Hash"
	"main/errores"
	user "main/usuarios"
	"strconv"
	"strings"
)

// Revisa si hay un usuario logueado
func chequear_existencia(User *string, err error) error {
	if *User == "" {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

// var Posts = TDAHash.CrearHash[int, user.Post]()       // [clave = ID: valor = POST]
// var Users = TDAHash.CrearHash[string, user.Usuario]() // [clave = nombre: valor = Usuario]

// Esta funcion lee todos los comandos que se utilizan durante la ejecucion del programa
func LecturaDeComandos(comando []string, Users TDAHash.Diccionario[string, user.Usuario], Posts TDAHash.Diccionario[int, user.Post], usuarioActual *string, ContadorPosts *int) {

	switch comando[0] {
	case "login":
		usuario := strings.Join(comando[1:], " ")

		//chequear que este loggueado
		if *usuarioActual != "" {
			fmt.Println(errores.ErrorUsuarioYaLoggeado{}.Error())
			return
		}

		//chequear que este en el dicc
		if !Users.Pertenece(usuario) {
			fmt.Println(errores.ErrorUsuarioNoExiste{}.Error())
			return
		}

		fmt.Printf("Hola %s\n", usuario)
		*usuarioActual = usuario

	case "logout":
		if chequear_existencia(usuarioActual, errores.ErrorUsuarioNoLoggeado{}) != nil {
			return
		}
		*usuarioActual = ""
		fmt.Println("Adios")

	case "publicar":
		if chequear_existencia(usuarioActual, errores.ErrorUsuarioNoLoggeado{}) != nil {
			return
		}
		texto := strings.Join(comando[1:], " ")
		usuario := Users.Obtener(*usuarioActual)
		usuario.Publicar(*ContadorPosts, texto, Posts, Users)
		*ContadorPosts++

	case "ver_siguiente_feed":
		if chequear_existencia(usuarioActual, errores.NoHayPostSiguiente{}) != nil {
			return
		}
		usuario := Users.Obtener(*usuarioActual)
		err := usuario.VerSiguientePost()
		if err != nil {
			fmt.Println(err.Error())
		}

	case "likear_post":
		if chequear_existencia(usuarioActual, errores.ErrorNoHayUsuarioOPostNoExiste{}) != nil {
			return
		}
		id, err1 := strconv.Atoi(comando[1])

		if err1 != nil {
			return
		}
		if !Posts.Pertenece(id) {
			fmt.Println(errores.ErrorNoHayUsuarioOPostNoExiste{}.Error())
			return
		}
		post := Posts.Obtener(id)
		post.RecibirLike(Users.Obtener(*usuarioActual))

	case "mostrar_likes":
		id, err1 := strconv.Atoi(comando[1])
		if err1 != nil {
			return
		}
		if !Posts.Pertenece(id) {
			fmt.Println(errores.ErrorPostNoExisteOSinLikes{}.Error())
			return
		}
		post := Posts.Obtener(id)
		err2 := post.MostrarLikes()
		if err2 != nil {
			fmt.Println(errores.ErrorPostNoExisteOSinLikes{}.Error())
			return
		}
	}
}

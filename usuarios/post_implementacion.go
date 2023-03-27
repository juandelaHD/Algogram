package usuarios

import (
	"fmt"
	TDAAbb "main/Abb"
	"main/errores"
	"strings"
)

type PostUsuario struct {
	texto   string
	likes   TDAAbb.DiccionarioOrdenado[string, string] // ABB de [nombre de usuario: nombre], compara en orden alfabetico
	id      int
	creador string
}

func CrearPost(id int, texto string, creador string) *PostUsuario {
	post := new(PostUsuario)
	post.id = id
	post.texto = texto
	post.likes = TDAAbb.CrearABB[string, string](strings.Compare)
	post.creador = creador
	return post
}

func (post *PostUsuario) RecibirLike(usuario Usuario) {
	fmt.Println("Post likeado")
	if post.likes.Pertenece(usuario.Nombre()) {
		return
	}
	post.likes.Guardar(usuario.Nombre(), usuario.Nombre())
	//Complejidad: O(log up)
}

func (post *PostUsuario) CantLikes() int {
	return post.likes.Cantidad()
}

func (post *PostUsuario) MostrarLikes() error {
	if post.CantLikes() == 0 {
		return errores.ErrorPostNoExisteOSinLikes{}
	}
	fmt.Printf("El post tiene %d likes:\n", post.CantLikes())
	post.likes.Iterar(func(clave string, _ string) bool {
		fmt.Printf("\t%s\n", clave)
		return true
	})
	return nil
	//Complejidad: O(up)
}

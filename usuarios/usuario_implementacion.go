package usuarios

import (
	"fmt"
	TDAHash "main/Hash"
	TDAHeap "main/Heap"
	"main/errores"
)

// Busca la distancia entre dos usuarios segun su posicion en el archivo, de la forma || pos1 - pos2 ||
func buscarDistancia(n1 int, n2 int) int {
	diferencia := n1 - n2
	if diferencia < 0 {
		return diferencia * -1
	}
	return diferencia
}

type Campo struct {
	afinidad int
	post     *PostUsuario
}

// Funcion de comparacion del heap. Primero compara por distancia (afinidad), y si es la misma, por id
func compararAfinidad(campo1 Campo, campo2 Campo) int {
	if campo1.afinidad < campo2.afinidad {
		return 1
	}
	if campo1.afinidad > campo2.afinidad {
		return -1
	}
	if campo1.post.id < campo2.post.id {
		return 1
	}
	return -1
}

func crearCampo(creador UsuarioRegistrado, receptor Usuario, post *PostUsuario) Campo {
	campo := new(Campo)
	campo.post = post
	campo.afinidad = buscarDistancia(creador.Index(), receptor.Index())
	return *campo
}

type UsuarioRegistrado struct {
	nombre string
	index  int
	feed   TDAHeap.ColaPrioridad[Campo]
}

func CrearUsuario(nombre string, index int) Usuario {
	new_user := new(UsuarioRegistrado)
	new_user.nombre = nombre
	new_user.index = index
	new_user.feed = TDAHeap.CrearHeap(compararAfinidad)

	return new_user
}

func (user *UsuarioRegistrado) Nombre() string {
	return user.nombre
}

func (user *UsuarioRegistrado) Index() int {
	return user.index
}

func (user *UsuarioRegistrado) Feed() TDAHeap.ColaPrioridad[Campo] {
	return user.feed
}

func (user *UsuarioRegistrado) Publicar(id int, texto string, diccPost TDAHash.Diccionario[int, Post], diccUsers TDAHash.Diccionario[string, Usuario]) {
	post := CrearPost(id, texto, user.nombre)
	diccPost.Guardar(id, post)
	iterador := diccUsers.Iterador()
	for iterador.HaySiguiente() {
		_, UsuarioActual := iterador.VerActual()
		if UsuarioActual.Nombre() != user.nombre {
			campo := crearCampo(*user, UsuarioActual, post)
			UsuarioActual.Feed().Encolar(campo)
		}
		iterador.Siguiente()
	}
	fmt.Println("Post publicado")
	//complejidad: O(u log(p))
}

func (user *UsuarioRegistrado) VerSiguientePost() error {
	if user.feed.EstaVacia() {
		return errores.NoHayPostSiguiente{}
	}
	campo := user.feed.Desencolar()
	fmt.Printf("Post ID %d\n", campo.post.id)
	fmt.Printf("%s dijo: %s\n", campo.post.creador, campo.post.texto)
	fmt.Printf("Likes: %d\n", campo.post.CantLikes())
	return nil
	//complejidad: O(log(p))
}

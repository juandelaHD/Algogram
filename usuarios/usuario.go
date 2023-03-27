package usuarios

import (
	TDAHash "main/Hash"
	TDAHeap "main/Heap"
)

type Usuario interface {
	//agarra el texto y lo coloca en el feed de cada usuario registrado
	//segun su afinidad
	Publicar(int, string, TDAHash.Diccionario[int, Post], TDAHash.Diccionario[string, Usuario])

	//Va al siguiente post del feed del usuario
	VerSiguientePost() error

	//Devuelve el nombre de Usuario
	Nombre() string

	//Devuelve el index del Usuario (el que tenia en el archivo)
	Index() int

	//Devuelve el feed del Usuario
	Feed() TDAHeap.ColaPrioridad[Campo]
}

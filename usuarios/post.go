package usuarios

type Post interface {
	//Recibir like hace que reciba el like de un usuario y coloca
	//su nombre de Usuario en el ABB si no se encuentra ya alli
	RecibirLike(usuario Usuario)

	//Devuelve el tamanio del ABB de usuarios que le dieron like
	CantLikes() int

	//Itera el ABB (si no esta vacio) e imprime todos los usuarios
	//que le dieron like en orden alfabetico
	MostrarLikes() error
}

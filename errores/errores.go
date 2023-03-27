package errores

type ErrorUsuarioYaLoggeado struct{}

func (e ErrorUsuarioYaLoggeado) Error() string {
	return "Error: Ya habia un usuario loggeado"
}

type ErrorUsuarioNoExiste struct{}

func (e ErrorUsuarioNoExiste) Error() string {
	return "Error: usuario no existente"
}

type ErrorUsuarioNoLoggeado struct{}

func (e ErrorUsuarioNoLoggeado) Error() string {
	return "Error: no habia usuario loggeado"
}

type NoHayPostSiguiente struct{}

func (e NoHayPostSiguiente) Error() string {
	return "Usuario no loggeado o no hay mas posts para ver"
}

type ErrorNoHayUsuarioOPostNoExiste struct{}

func (e ErrorNoHayUsuarioOPostNoExiste) Error() string {
	return "Error: Usuario no loggeado o Post inexistente"
}

type ErrorPostNoExisteOSinLikes struct{}

func (e ErrorPostNoExisteOSinLikes) Error() string {
	return "Error: Post inexistente o sin likes"
}

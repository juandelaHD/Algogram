package lista

type Lista[T any] interface {

	// EstaVacia devuelve verdadero si la lista no tiene elementos, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero agrega el elemento a la primera posicion de la lista.
	InsertarPrimero(elem T)

	// InsertarUltimo agrega el elemento a la ultima posicion de la lista.
	InsertarUltimo(elem T)

	// BorrarPrimero elimina el primer elemento de la lista. Si esta vacia, entra en panico con un mensaje
	// "La lista esta vacia".
	BorrarPrimero() T

	// VerPrimero obtiene el valor del primer elemento de la lista. Si esta vacia, entra en panico con un mensaje
	// "La lista esta vacia".
	VerPrimero() T

	// VerUltimo obtiene el valor del ultimo dato de la lista. Si esta vacia, entra en panico con un mensaje
	// "La lista esta vacia".
	VerUltimo() T

	// Largo devuelve la cantidad de elementos de la lista.
	Largo() int

	// Iterar recorre a cada uno de los datos de la lista (de primero al ultimo), hasta que la lista se termine
	// o la funcion visitar devuelva false.
	Iterar(visitar func(T) bool)

	// Iterador devuelve un Iterador que permite recorrer cada dato de la lista.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	// VerActual obtiene el valor del elemento de la lista que apunta el iterador. Si esta vacia, entra en panico con un mensaje
	// "La lista esta vacia".
	VerActual() T

	// HaySiguiente devuelve verdadero si hay proximo elemento para iterar, false en caso contrario.
	HaySiguiente() bool

	// Siguiente devuelve el proximo elemento a la iteracion.
	Siguiente() T

	// Insertar agrega un elemento en donde se encuentre el iterador en la lista.
	Insertar(T)

	// Borrar elimina el elemento que se esta iterando y lo devuelve.
	Borrar() T
}

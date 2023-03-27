package lista

type nodoLista[T any] struct {
	proximo *nodoLista[T]
	dato    T
}

type ListaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iterListaEnlazada[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	lista    *ListaEnlazada[T]
}

func crearNodo[T any](dato T) *nodoLista[T] {
	nodo := new(nodoLista[T])
	nodo.dato = dato
	nodo.proximo = nil
	return nodo
}

func CrearListaEnlazada[T any]() Lista[T] {
	return new(ListaEnlazada[T])
}

func (lista ListaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista ListaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista ListaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.primero.dato
}

func (lista ListaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.ultimo.dato
}

func (lista *ListaEnlazada[T]) InsertarPrimero(elem T) {
	agregado := crearNodo(elem)
	agregado.proximo = lista.primero
	if lista.EstaVacia() {
		lista.ultimo = agregado
	}
	lista.primero = agregado
	lista.largo++
}

func (lista *ListaEnlazada[T]) InsertarUltimo(elem T) {
	agregado := crearNodo(elem)

	if lista.EstaVacia() {
		lista.primero = agregado
	} else {
		actual := lista.primero
		for actual.proximo != nil {
			actual = actual.proximo
		}
		actual.proximo = agregado
	}

	lista.ultimo = agregado
	lista.largo++
}

func (lista *ListaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	elem := lista.VerPrimero()
	lista.primero = lista.primero.proximo
	if lista.Largo() == 1 {
		lista.ultimo = nil
	}
	lista.largo--
	return elem
}

func (lista *ListaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := lista.primero
	for actual != nil && visitar(actual.dato) {
		actual = actual.proximo
	}
}

func (lista *ListaEnlazada[T]) Iterador() IteradorLista[T] {
	iter := new(iterListaEnlazada[T])
	iter.lista = lista
	iter.actual = lista.primero
	return iter
}

func (iter *iterListaEnlazada[T]) VerActual() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.actual.dato
}

func (iter *iterListaEnlazada[T]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iterListaEnlazada[T]) Siguiente() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	elem := iter.actual.dato
	iter.anterior = iter.actual
	iter.actual = iter.actual.proximo
	return elem
}

func (iter *iterListaEnlazada[T]) Borrar() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	elem := iter.actual.dato
	if iter.lista.largo == 1 {
		iter.actual = nil
		iter.lista.primero = nil
		iter.lista.ultimo = nil

	} else if iter.anterior == nil {
		iter.lista.primero = iter.actual.proximo
		iter.actual = iter.actual.proximo

	} else if iter.actual.proximo == nil {
		iter.actual = nil
		iter.anterior.proximo = nil
		iter.lista.ultimo = iter.anterior

	} else {
		iter.anterior.proximo = iter.actual.proximo
		iter.actual = iter.anterior.proximo
	}
	iter.lista.largo--
	return elem
}

func (iter *iterListaEnlazada[T]) Insertar(elem T) {
	agregado := crearNodo(elem)
	agregado.proximo = iter.actual

	if iter.anterior == nil {
		iter.lista.primero = agregado
		if iter.actual == nil {
			iter.lista.ultimo = agregado
		}

	} else if iter.actual == nil {
		iter.lista.ultimo = agregado
		iter.anterior.proximo = agregado

	} else {
		iter.anterior.proximo = agregado
	}

	iter.actual = agregado
	iter.lista.largo++
}

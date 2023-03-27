package diccionario

import (
	TDAPila "main/pila"
)

type nodoAbb[K comparable, V any] struct {
	izq   *nodoAbb[K, V]
	der   *nodoAbb[K, V]
	clave K
	valor V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      func(K, K) int
}

type iterDiccionario[K comparable, V any] struct {
	stack TDAPila.Pila[*nodoAbb[K, V]]
}

type iterDiccionarioRango[K comparable, V any] struct {
	stack TDAPila.Pila[*nodoAbb[K, V]]
	cmp   func(K, K) int
	desde *K
	hasta *K
}

// Creaciones de nodos e interfaces ------------------------------------------------------------------

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	abb := new(abb[K, V])
	abb.cmp = funcion_cmp
	return abb
}

func crearNodoArbol[K comparable, V any](clave K, valor V) *nodoAbb[K, V] {
	nodo := new(nodoAbb[K, V])
	nodo.clave = clave
	nodo.valor = valor
	return nodo
}

// Metodos del DiccionarioOrdenado ------------------------------------------------------------------

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

// --------------------------------------------------------------------------------------------------

func (abb *abb[K, V]) Obtener(clave K) V {
	return obtener(abb.raiz, clave, abb.cmp)
}

func obtener[K comparable, V any](actual *nodoAbb[K, V], clave K, comparar func(K, K) int) V {
	if actual == nil {
		panic("La clave no pertenece al diccionario")
	}
	comparacion := comparar(actual.clave, clave)
	if comparacion == 0 {
		return actual.valor
	}
	if comparacion < 0 {
		return obtener(actual.der, clave, comparar)
	}
	return obtener(actual.izq, clave, comparar)
}

// --------------------------------------------------------------------------------------------------

func (abb *abb[K, V]) Pertenece(clave K) bool {
	return pertenece(abb.raiz, clave, abb.cmp)
}

func pertenece[K comparable, V any](actual *nodoAbb[K, V], clave K, comparar func(K, K) int) bool {
	if actual == nil {
		return false
	}
	comparacion := comparar(actual.clave, clave)
	if comparacion == 0 {
		return true
	}
	if comparacion < 0 {
		return pertenece(actual.der, clave, comparar)
	}
	return pertenece(actual.izq, clave, comparar)
}

// --------------------------------------------------------------------------------------------------
func (abb *abb[K, V]) Guardar(clave K, valor V) {
	nuevoNodo := crearNodoArbol(clave, valor)
	if abb.raiz == nil {
		abb.raiz = nuevoNodo
		abb.cantidad++
	} else {
		cantidadASumar := guardar(abb.raiz, abb.cmp, nuevoNodo, nil)
		abb.cantidad += cantidadASumar
	}
}

func guardar[K comparable, V any](actual *nodoAbb[K, V], comparar func(K, K) int, nuevoNodo *nodoAbb[K, V], anterior *nodoAbb[K, V]) int {
	if actual == nil {
		comparacion := comparar(anterior.clave, nuevoNodo.clave)
		if comparacion > 0 {
			anterior.izq = nuevoNodo
		} else {
			anterior.der = nuevoNodo
		}
		return 1
	}
	comparacion := comparar(actual.clave, nuevoNodo.clave)
	if comparacion == 0 {
		actual.valor = nuevoNodo.valor
		return 0
	}
	if comparacion > 0 {
		return guardar(actual.izq, comparar, nuevoNodo, actual)
	} else {
		return guardar(actual.der, comparar, nuevoNodo, actual)
	}
}

// --------------------------------------------------------------------------------------------------

func (abb *abb[K, V]) Borrar(clave K) V {
	var borrado V
	actual := abb.raiz
	if actual != nil && actual.clave == clave {
		borrado = actual.valor
		borrarRaiz(abb, actual.der, actual.izq)
	} else {
		borrado = borrar(actual, clave, abb.cmp, nil)
	}
	abb.cantidad--
	return borrado
}

func borrar[K comparable, V any](actual *nodoAbb[K, V], clave K, comparar func(K, K) int, anterior *nodoAbb[K, V]) V {
	if actual == nil {
		panic("La clave no pertenece al diccionario")
	}
	comparacion := comparar(actual.clave, clave)
	if comparacion == 0 {
		return borrarNodo(actual, anterior)
	}
	if comparacion > 0 {
		return borrar(actual.izq, clave, comparar, actual)
	}
	return borrar(actual.der, clave, comparar, actual)
}

func borrarNodo[K comparable, V any](actual *nodoAbb[K, V], anterior *nodoAbb[K, V]) V {
	valor := actual.valor
	if actual.izq == nil && actual.der == nil {
		if anterior.izq != nil && anterior.izq.clave == actual.clave {
			anterior.izq = nil
		} else {
			anterior.der = nil
		}
	} else if actual.izq == nil {
		if anterior.izq != nil && anterior.izq.clave == actual.clave {
			anterior.izq = actual.der
		} else {
			anterior.der = actual.der
		}
	} else if actual.der == nil {
		if anterior.izq != nil && anterior.izq.clave == actual.clave {
			anterior.izq = actual.izq
		} else {
			anterior.der = actual.izq
		}
	} else {
		reemplazante := buscarReemplazante(actual.der)
		reemplazante.der = borrarReemplazo(actual.der)
		reemplazante.izq = actual.izq
		if anterior.izq != nil && anterior.izq.clave == actual.clave {
			anterior.izq = reemplazante
		} else {
			anterior.der = reemplazante
		}
	}
	return valor
}

func buscarReemplazante[K comparable, V any](actual *nodoAbb[K, V]) *nodoAbb[K, V] {
	if actual.izq == nil {
		return actual
	}
	return buscarReemplazante(actual.izq)
}

func borrarReemplazo[K comparable, V any](actual *nodoAbb[K, V]) *nodoAbb[K, V] {
	if actual.izq == nil {
		return actual.der
	}
	actual.izq = borrarReemplazo(actual.izq)
	return actual
}

func borrarRaiz[K comparable, V any](abb *abb[K, V], hijoDer *nodoAbb[K, V], hijoIzq *nodoAbb[K, V]) {
	if hijoIzq == nil && hijoDer == nil {
		abb.raiz = nil
	} else if hijoIzq != nil && hijoDer != nil {
		abb.raiz = hijoDer
		reemplazante := hijoDer
		for reemplazante.izq != nil {
			reemplazante = reemplazante.izq
		}
		reemplazante.izq = hijoIzq
	} else if hijoIzq == nil {
		abb.raiz = hijoDer
	} else {
		abb.raiz = hijoIzq
	}
}

// --------------------------------------------------------------------------------------------------

func (abb *abb[K, V]) Iterar(visitar func(clave K, valor V) bool) {
	iterar(visitar, abb.raiz)
}

func iterar[K comparable, V any](visitar func(clave K, valor V) bool, actual *nodoAbb[K, V]) {
	if actual == nil {
		return
	}
	iterar(visitar, actual.izq)
	if !visitar(actual.clave, actual.valor) {
		return
	}
	iterar(visitar, actual.der)
}

// --------------------------------------------------------------------------------------------------

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(iterDiccionario[K, V])
	iter.stack = TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	actual := abb.raiz
	for actual != nil {
		iter.stack.Apilar(actual)
		actual = actual.izq
	}
	return iter
}

// Metodos del Iterador -----------------------------------------------------------------------------

func (iter *iterDiccionario[K, V]) HaySiguiente() bool {
	return !iter.stack.EstaVacia()
}

// --------------------------------------------------------------------------------------------------

func (iter *iterDiccionario[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	tope := iter.stack.VerTope()
	return tope.clave, tope.valor
}

// --------------------------------------------------------------------------------------------------

func (iter *iterDiccionario[K, V]) Siguiente() K {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := iter.stack.Desapilar()
	actual := nodo.der
	for actual != nil {
		iter.stack.Apilar(actual)
		actual = actual.izq
	}
	return nodo.clave
}

// Iterador por Rango -------------------------------------------------------------------------------

func abbMinimo[K comparable, V any](actual *nodoAbb[K, V]) *nodoAbb[K, V] {
	if actual == nil {
		return nil
	}
	for actual.izq != nil {
		actual = actual.izq
	}
	return actual
}

func abbMaximo[K comparable, V any](actual *nodoAbb[K, V]) *nodoAbb[K, V] {
	if actual == nil {
		return nil
	}
	for actual.der != nil {
		actual = actual.der
	}
	return actual
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	if desde == nil {
		desde = &abbMinimo(abb.raiz).clave
	}
	if hasta == nil {
		hasta = &abbMaximo(abb.raiz).clave
	}
	iterarRango(visitar, abb.raiz, desde, hasta, abb.cmp)
}

func iterarRango[K comparable, V any](visitar func(clave K, valor V) bool, actual *nodoAbb[K, V], desde *K, hasta *K, comparar func(K, K) int) {
	if actual == nil {
		return
	}
	iterarRango(visitar, actual.izq, desde, hasta, comparar)
	if comparar(actual.clave, *desde) >= 0 && comparar(actual.clave, *hasta) <= 0 {
		if !visitar(actual.clave, actual.valor) {
			return
		}
	}
	iterarRango(visitar, actual.der, desde, hasta, comparar)
}

// --------------------------------------------------------------------------------------------------

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	if desde == nil {
		desde = &abbMinimo(abb.raiz).clave
	}
	if hasta == nil {
		hasta = &abbMaximo(abb.raiz).clave
	}
	iter := new(iterDiccionarioRango[K, V])
	iter.stack = TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iter.cmp = abb.cmp
	iter.desde = desde
	iter.hasta = hasta

	actual := abb.raiz
	for actual != nil && abb.cmp(actual.clave, *desde) < 0 {
		actual = actual.der
	}
	for actual != nil {
		if abb.cmp(actual.clave, *desde) >= 0 {
			iter.stack.Apilar(actual)
			actual = actual.izq
		} else {
			actual = actual.der
		}
	}
	return iter
}

// Metodos del Iterador por Rango --------------------------------------------------------------------

func (iter *iterDiccionarioRango[K, V]) HaySiguiente() bool {
	return !iter.stack.EstaVacia() && iter.cmp(iter.stack.VerTope().clave, *iter.hasta) <= 0
}

// --------------------------------------------------------------------------------------------------

func (iter *iterDiccionarioRango[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := iter.stack.VerTope()
	return nodo.clave, nodo.valor
}

// --------------------------------------------------------------------------------------------------

func (iter *iterDiccionarioRango[K, V]) Siguiente() K {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := iter.stack.Desapilar()
	clave := nodo.clave
	actual := nodo.der
	for actual != nil && iter.cmp(actual.clave, *iter.desde) >= 0 {
		iter.stack.Apilar(actual)
		actual = actual.izq
	}
	return clave
}

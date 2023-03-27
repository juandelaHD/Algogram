package pila

const (
	_DOS      int    = 2
	_CUARTO   int    = 4
	_CAPMAX   int    = 10
	_MSGERROR string = "La pila esta vacia"
)

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, _CAPMAX)
	return pila
}

func (pila *pilaDinamica[T]) redimensionar(n int) {
	NuevaPila := make([]T, n)
	copy(NuevaPila, pila.datos)
	pila.datos = NuevaPila
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	return (pila.cantidad) == 0
}

func (pila pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic(_MSGERROR)
	}
	elem := pila.datos[(pila.cantidad - 1)]
	return elem
}

func (pila *pilaDinamica[T]) Apilar(elem T) {
	if pila.cantidad == cap(pila.datos) {
		pila.redimensionar(cap(pila.datos) * _DOS)
	}
	(pila.datos[pila.cantidad]) = elem
	(pila.cantidad)++
}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic(_MSGERROR)
	}
	elem := (pila.datos[pila.cantidad-1])
	(pila.cantidad)--
	if pila.cantidad <= cap(pila.datos)/_CUARTO && cap(pila.datos) > _CAPMAX {
		pila.redimensionar(cap(pila.datos) / _DOS)
	}
	return elem
}

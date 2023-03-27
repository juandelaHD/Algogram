package cola_prioridad

func _Swap[T comparable](x *T, y *T) {
	*x, *y = *y, *x
}

type heap[T comparable] struct {
	datos    []T
	cmp      func(T, T) int
	cantidad int
}

func _Upheap[T comparable](arr []T, pos int, cmp func(T, T) int) {
	if pos == 0 {
		return
	}
	padre := (pos - 1) / 2
	if cmp(arr[padre], arr[pos]) < 0 {
		_Swap(&arr[pos], &arr[padre])
		_Upheap(arr, padre, cmp)
	}
}

func _Downheap[T comparable](arr []T, pos int, cmp func(T, T) int) {
	hijo_izq := (pos * 2) + 1
	hijo_der := (pos * 2) + 2
	if hijo_izq >= len(arr) {
		return
	}
	if hijo_der >= len(arr) {
		if cmp(arr[pos], arr[hijo_izq]) < 0 {
			_Swap(&arr[pos], &arr[hijo_izq])
		}
		return
	}
	if cmp(arr[hijo_izq], arr[hijo_der]) > 0 {
		if cmp(arr[hijo_izq], arr[pos]) > 0 {
			_Swap(&arr[hijo_izq], &arr[pos])
			_Downheap(arr, hijo_izq, cmp)
		}
	} else {
		if cmp(arr[hijo_der], arr[pos]) > 0 {
			_Swap(&arr[hijo_der], &arr[pos])
			_Downheap(arr, hijo_der, cmp)
		}
	}
}

func _Heapify[T comparable](arr []T, cmp func(T, T) int) []T {
	for i := len(arr) - 1; i >= 0; i-- {
		_Downheap(arr, i, cmp)
	}
	return arr
}

func CrearHeap[T comparable](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heap[T])
	heap.datos = []T{}
	heap.cmp = funcion_cmp
	return heap
}

func CrearHeapArr[T comparable](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	final := make([]T, len(arreglo))
	copy(final, arreglo)
	heap := new(heap[T])
	heap.cmp = funcion_cmp
	heap.cantidad = len(arreglo)
	heap.datos = _Heapify(final, heap.cmp)
	return heap
}

func (heap heap[T]) EstaVacia() bool {
	return heap.cantidad == 0
}

func (heap heap[T]) Cantidad() int {
	return heap.cantidad
}

func (heap heap[T]) VerMax() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	return heap.datos[0]
}

func (heap *heap[T]) Encolar(elem T) {
	heap.datos = append(heap.datos, elem)
	_Upheap(heap.datos, heap.cantidad, heap.cmp)
	heap.cantidad++
}

func (heap *heap[T]) Desencolar() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	primero := heap.datos[0]
	heap.cantidad--
	_Swap(&heap.datos[0], &heap.datos[heap.cantidad])
	heap.datos = heap.datos[:heap.cantidad]
	_Downheap(heap.datos, 0, heap.cmp)
	return primero
}

func HeapSort[T comparable](elementos []T, funcion_cmp func(T, T) int) {
	elementos = _Heapify(elementos, funcion_cmp)
	for i := len(elementos); i > 0; i-- {
		_Downheap(elementos[:i], 0, funcion_cmp)
		_Swap(&elementos[0], &elementos[i-1])
	}
}

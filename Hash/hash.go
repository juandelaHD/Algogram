package diccionario

import (
	"fmt"
	"hash/fnv"
	TDALista "main/lista"
)

const _TAM_DICCIONARIO_INICIAL = 101
const _FACTOR_DE_CARGA = 3

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func (dicc *diccionarioHash[K, V]) hash(clave K) uint32 {
	h := fnv.New32a()
	h.Write(convertirABytes(clave))
	return h.Sum32() % dicc.capacidad
}

func es_primo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func buscarPrimoCercano(capacidad int) int {
	for i := capacidad; i < capacidad*_FACTOR_DE_CARGA; i++ {
		if es_primo(i) {
			return i
		}
	}
	return capacidad*_FACTOR_DE_CARGA + 1
}

// STRUCT DEL DICCIONARIO + ITER + CAMPO

type campo[K comparable, V any] struct {
	clave K
	valor V
}

type diccionarioHash[K comparable, V any] struct {
	arreglo   []TDALista.Lista[campo[K, V]]
	cantidad  int
	capacidad uint32
}

type iterDiccionarioHash[K comparable, V any] struct {
	diccionario      *diccionarioHash[K, V]
	actual           TDALista.IteradorLista[campo[K, V]]
	contadorDeCampos int
	index            int
}

// crear HASH y crear CAMPO

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	arr := make([]TDALista.Lista[campo[K, V]], _TAM_DICCIONARIO_INICIAL)
	for index := range arr {
		arr[index] = TDALista.CrearListaEnlazada[campo[K, V]]()
	}
	dicc := new(diccionarioHash[K, V])
	dicc.arreglo = arr
	dicc.capacidad = _TAM_DICCIONARIO_INICIAL
	return dicc
}

func crearcampo[K comparable, V any](clave K, valor V) campo[K, V] {
	nuevocampo := new(campo[K, V])
	nuevocampo.clave = clave
	nuevocampo.valor = valor
	return *nuevocampo
}

// METODOS DE DICCIONARIO

func (dicc *diccionarioHash[K, V]) redimensionar(nuevoCapacidad int) {
	guardado := make([]campo[K, V], dicc.Cantidad())
	var contador int = 0

	for _, lista_enlazada := range dicc.arreglo {
		iterador := lista_enlazada.Iterador()
		for iterador.HaySiguiente() {
			guardado[contador] = iterador.VerActual()
			iterador.Siguiente()
			contador++
		}
	}

	arregloRedimensionado := make([]TDALista.Lista[campo[K, V]], nuevoCapacidad)
	for index := range arregloRedimensionado {
		arregloRedimensionado[index] = TDALista.CrearListaEnlazada[campo[K, V]]()
	}
	dicc.arreglo = arregloRedimensionado
	dicc.capacidad = uint32(nuevoCapacidad)
	dicc.cantidad = 0

	for _, campos := range guardado {
		dicc.Guardar(campos.clave, campos.valor)
	}
}

func (dicc diccionarioHash[K, V]) Cantidad() int {
	return dicc.cantidad
}

func (dicc *diccionarioHash[K, V]) Guardar(clave K, valor V) {
	if dicc.Cantidad() >= _FACTOR_DE_CARGA*int(dicc.capacidad) {
		dicc.redimensionar(buscarPrimoCercano(int(dicc.capacidad) * _FACTOR_DE_CARGA))
	}

	posicion := dicc.hash(clave)
	dato := crearcampo(clave, valor)
	if dicc.Pertenece(clave) {
		dicc.Borrar(clave)
	}
	dicc.cantidad++
	dicc.arreglo[posicion].InsertarUltimo(dato)
}

func (dicc diccionarioHash[K, V]) Pertenece(clave K) bool {
	posicion := dicc.hash(clave)
	iterador := dicc.arreglo[posicion].Iterador()
	for iterador.HaySiguiente() {
		if iterador.VerActual().clave == clave {
			return true
		}
		iterador.Siguiente()
	}
	return false
}

func (dicc diccionarioHash[K, V]) Obtener(clave K) V {
	posicion := dicc.hash(clave)
	iterador := dicc.arreglo[posicion].Iterador()
	for iterador.HaySiguiente() {
		if iterador.VerActual().clave == clave {
			return iterador.VerActual().valor
		}
		iterador.Siguiente()
	}
	panic("La clave no pertenece al diccionario")
}

func (dicc *diccionarioHash[K, V]) Borrar(clave K) V {
	if dicc.capacidad > uint32(_TAM_DICCIONARIO_INICIAL) && dicc.Cantidad() <= _FACTOR_DE_CARGA/int(dicc.capacidad) {
		dicc.redimensionar(buscarPrimoCercano(int(dicc.capacidad) / _FACTOR_DE_CARGA))
	}

	posicion := dicc.hash(clave)
	iterador := dicc.arreglo[posicion].Iterador()
	for iterador.HaySiguiente() {
		if iterador.VerActual().clave == clave {
			dicc.cantidad--
			return iterador.Borrar().valor
		}
		iterador.Siguiente()
	}
	panic("La clave no pertenece al diccionario")
}

func (dicc *diccionarioHash[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for _, lista_enlazada := range dicc.arreglo {
		iter := lista_enlazada.Iterador()
		for iter.HaySiguiente() {
			if !visitar(iter.VerActual().clave, iter.VerActual().valor) {
				return
			}
			iter.Siguiente()
		}
	}
}

func (dicc *diccionarioHash[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(iterDiccionarioHash[K, V])
	iter.diccionario = dicc
	for _, listaEnlazada := range iter.diccionario.arreglo {
		if !listaEnlazada.EstaVacia() {
			iter.actual = listaEnlazada.Iterador()
			iter.contadorDeCampos = 1
			break
		}
		iter.index++
	}
	return iter

}

// METODOS ITERADOR

func (iter *iterDiccionarioHash[K, V]) HaySiguiente() bool {
	return iter.contadorDeCampos <= iter.diccionario.Cantidad() && iter.actual != nil
}

func (iter *iterDiccionarioHash[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.actual.VerActual().clave, iter.actual.VerActual().valor
}

func (iter *iterDiccionarioHash[K, V]) Siguiente() K {
	clave, _ := iter.VerActual()
	iter.actual.Siguiente()
	if !iter.actual.HaySiguiente() {
		for _, listaEnlazada := range iter.diccionario.arreglo[iter.index+1:] {
			iter.index++
			if listaEnlazada.EstaVacia() {
				iter.actual = nil
			} else {
				iter.actual = listaEnlazada.Iterador()
				break
			}
		}
	}
	iter.contadorDeCampos++
	return clave
}

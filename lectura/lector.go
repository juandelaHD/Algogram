package lectura

import (
	"bufio"
	"os"
)

//leemos el archivo y colocamos los nombres en una lista

func LeerUsuarios(ruta string) []string {
	listaUsuarios := []string{}

	archivo, err := os.Open(ruta)
	if err != nil {
		return nil
	}
	defer archivo.Close()
	lector := bufio.NewScanner(archivo)

	for lector.Scan() {
		listaUsuarios = append(listaUsuarios, lector.Text())
	}
	err = lector.Err()
	if err != nil {
		return nil
	}

	return listaUsuarios
}

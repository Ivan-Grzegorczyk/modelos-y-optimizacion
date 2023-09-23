package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

/* Acá pongo todas las funciones relacionadas con leer los datos del archivo */

func abrirArchivo(s string) *os.File {
	archivo, err := os.Open(s)
	if err != nil {
		panic(err)
	}
	return archivo
}

func cargarDatos(sucursales *[]sucursal, capacidad *int, dimension *int) {
	archivo := abrirArchivo("cmd/resources/data.txt")
	defer func(archivo *os.File) {
		err := archivo.Close()
		if err != nil {
			panic(err)
		}
	}(archivo)

	scanner := bufio.NewScanner(archivo)
	scanner.Scan()
	if scanner.Text()[0:10] != "CAPACIDAD:" {
		panic("Error en formato de archivo")
	}
	*capacidad, _ = strconv.Atoi(scanner.Text()[11:])
	scanner.Scan()
	if scanner.Text()[0:10] != "DIMENSION:" {
		panic("Error en formato de archivo")
	}
	*dimension, _ = strconv.Atoi(scanner.Text()[11:])
	moverseHasta(scanner, "DEMANDAS")
	for i := 1; i <= *dimension; i++ {
		linea := scanner.Text()
		demandaNodo, _ := strconv.Atoi(strings.Split(linea, " ")[1])
		*sucursales = append(*sucursales, sucursal{
			id:       i,
			visitada: false,
			demanda:  demandaNodo,
		})
		scanner.Scan()
	}
	moverseHasta(scanner, "NODE_COORD_SECTION")
	for i := 1; i <= *dimension; i++ {
		linea := scanner.Text()
		x, _ := strconv.ParseFloat(strings.Split(linea, " ")[1], 64)
		y, _ := strconv.ParseFloat(strings.Split(linea, " ")[2], 64)
		(*sucursales)[i-1].x = x
		(*sucursales)[i-1].y = y
		scanner.Scan()
	}
}

func moverseHasta(scanner *bufio.Scanner, line string) {
	for scanner.Text() != line {
		scanner.Scan()
	}
	scanner.Scan()
}

func imprimirResultado(recorrido []sucursal) {
	archivo, err := os.Create("entrega_primer_problema.txt")
	if err != nil {
		panic(err)
	}
	defer func(archivo *os.File) {
		err := archivo.Close()
		if err != nil {
			panic(err)
		}
	}(archivo)

	for _, sucursal := range recorrido {
		_, err := archivo.WriteString(strconv.Itoa(sucursal.id) + " ")
		if err != nil {
			panic(err)
		}
	}
}

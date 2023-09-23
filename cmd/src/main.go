package main

import (
	"fmt"
	"math"
)

type sucursal struct {
	id       int
	x, y     float64
	demanda  int
	visitada bool
}

func main() {
	var sucursales []sucursal
	var capacidad, dimension int

	cargarDatos(&sucursales, &capacidad, &dimension)

	// Busco la sucursal más cercana a la actual que cumpla la capacidad y me muevo a la misma eliminando la anterior
	n := len(sucursales)
	visitadas := make([]bool, n)
	var recorrido []sucursal

	indiceActual := 0
	sucursalActual := sucursales[indiceActual]
	cargaActual := 0
	recorrido = append(recorrido, sucursalActual)
	visitadas[indiceActual] = true

	distanciaTotal := float64(0)

	for len(recorrido) < n {
		indiceSiguiente, sucursalSiguiente := calcularSiguiente(sucursalActual, sucursales, visitadas, cargaActual, capacidad)
		indiceActual = indiceSiguiente
		distanciaTotal += calcularDistancia(sucursalActual, sucursalSiguiente)
		sucursalActual = sucursalSiguiente
		cargaActual += sucursalActual.demanda
		recorrido = append(recorrido, sucursalActual)
		visitadas[indiceActual] = true
	}

	// Imprimo el recorrido
	fmt.Println("El largo del recorrido es:", len(recorrido))
	fmt.Println("La distancia total es:", distanciaTotal)
	fmt.Println("******* Recorrido ********")
	for _, sucursal := range recorrido {
		print(sucursal.id, " ")
	}
}

func calcularDistancia(sucursal1, sucursal2 sucursal) float64 {
	return math.Sqrt(math.Pow(sucursal2.x-sucursal1.x, 2) + math.Pow(sucursal2.y-sucursal1.y, 2))
}

func calcularSiguiente(sucursalActual sucursal, sucursales []sucursal, visitadas []bool, cargaActual, capacidad int) (int, sucursal) {
	distanciaMinima := math.Inf(1) // Inicializo en infinito
	var IndiceSucursalCercana int
	var sucursalCercana sucursal

	for i, sucursal := range sucursales {
		if !visitadas[i] && cargaActual+sucursal.demanda <= capacidad && cargaActual+sucursal.demanda > 0 {
			distancia := calcularDistancia(sucursalActual, sucursal)
			if distancia < distanciaMinima {
				distanciaMinima = distancia
				IndiceSucursalCercana = i
				sucursalCercana = sucursal
			}
		}
	}

	return IndiceSucursalCercana, sucursalCercana
}

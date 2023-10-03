package main

import (
	"fmt"
	"math"
	"time"
)

type sucursal struct {
	id      int
	x, y    float64
	demanda int
}

func main() {
	var sucursales []sucursal
	var capacidad int
	// primer, segundo
	nroProblema := "segundo"

	cargarDatos(nroProblema+"_problema.txt", &sucursales, &capacidad)

	// Busco la sucursal más cercana a la actual que cumpla la capacidad y me muevo a la misma eliminando la anterior
	dimension := len(sucursales)
	// Cantidad de combinaciones a probar
	profundidad := 100 // Debe ser 1 ≤ profundidad ≤ dimension

	recorrido := resolverProblema(profundidad, sucursales, dimension, capacidad)

	// Imprimo el recorrido
	imprimirResultado(nroProblema, recorrido) // Escribo el resultado en un archivo
}

func resolverProblema(profundidad int, sucursales []sucursal, dimension int, capacidad int) []sucursal {
	distanciaMejorRecorrido := math.Inf(1)
	var recorrido []sucursal

	start := time.Now()
	for i := 0; i < profundidad; i++ {
		var distancia float64
		var recorridoActual []sucursal
		if sucursales[i].demanda >= 0 {
			distancia, recorridoActual = calcularRecorrido(i, dimension, sucursales, capacidad)
			if distancia < distanciaMejorRecorrido {
				distanciaMejorRecorrido = distancia
				recorrido = recorridoActual
			}
		}
	}
	elapsed := time.Since(start)

	// Tiempo que tarda en ejecutarse
	fmt.Printf("El programa tardó %s en ejecutarse\n", elapsed)
	fmt.Printf("La distancia total es: %.2f", distanciaMejorRecorrido)

	return recorrido
}

func calcularRecorrido(indiceInicial int, n int, sucursales []sucursal, capacidad int) (float64, []sucursal) {
	visitadas := make([]bool, n)
	distanciaTotal := float64(0)
	var recorrido []sucursal

	indiceActual := indiceInicial
	sucursalActual := sucursales[indiceActual]
	cargaActual := sucursales[indiceActual].demanda
	recorrido = append(recorrido, sucursalActual)
	visitadas[indiceActual] = true

	for len(recorrido) < n {
		indiceSiguiente, sucursalSiguiente := calcularSiguiente(sucursalActual, sucursales, visitadas, cargaActual, capacidad)
		indiceActual = indiceSiguiente
		distanciaTotal += calcularDistancia(sucursalActual, sucursalSiguiente)
		sucursalActual = sucursalSiguiente
		cargaActual += sucursalActual.demanda
		recorrido = append(recorrido, sucursalActual)
		visitadas[indiceActual] = true
	}
	distanciaTotal += calcularDistancia(sucursalActual, sucursales[indiceInicial])
	return distanciaTotal, recorrido
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

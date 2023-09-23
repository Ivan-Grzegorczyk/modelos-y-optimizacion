# Trabajo Práctico 1 Modelos y Optimización

### Impresiones del problema
El problema me pareció de complejidad intermedia, es difícil pensar como 
empezar a plantearlo, pero una vez que se entiende la idea es bastante sencillo.
Elegí este lenguaje por encima de otros más sencillos como Python principalmente
porque quería aprender a usarlo y porque sé que es mucho más rápido, teniendo en 
cuenta la naturaleza del problema pensé que la velocidad era un factor importante.
El problema me pareció muy similar al "Traveling Salesman Problem", pero más complejo
porque incluye el factor de la capacidad máxima del camión.

### Ideas para resolver el problema
La idea es recorrer la lista de sucursales una a una buscando la más cercana a la
actual que cumpla con la condición de impuesta por el problema sobre la cantidad de 
dinero a transportar. Basándome en esto voy armando el recorrido.
En una próxima entrega voy a intentar optimizar más el resultado calculando el
recorrido iniciando desde cualquier punto que cargue dinero en el camión.
Se podría mejorar más el resultado si a la hora de ir calculando el recorrido se
usara más profundidad para elegir la próxima sucursal a visitar, pero eso aumentaría
mucho la complejidad del algoritmo, puede que en una próxima entrega lo intente encarar.

### Comentarios sobre el código
* Primer entrega: 
    Inicialmente, quise plantearlo armando un grafo y poniendo las distancias como los
    pesos de las aristas, pero se me complicó mucho implementarlo, por lo que decidí
    usar slices de Go. Se inicia de forma arbitraria por la primera sucursal planteada
    en el archivo de texto. Distancia total recorrida: 8691.176773576135
* Segunda entrega:
    En lugar de iniciar el recorrido desde el primer punto del archivo, se calcula
    cuál es más óptimo para el algoritmo actual. Distancia total recorrida: 8032.118854705603

## Comentarios finales
El problema me pareció muy interesante, me gustó mucho resolverlo. Es el primer
problema de optimización que resuelvo, y me gustó mucho la experiencia.

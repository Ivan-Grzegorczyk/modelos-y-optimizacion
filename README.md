# Trabajo Práctico 1 Modelos y Optimización

## Problema 1

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
  en el archivo de texto.
  Distancia total recorrida: 9082.86
  Tiempo de ejecución: 498.1 µs
* Segunda entrega:
  En lugar de iniciar el recorrido desde el primer punto del archivo, se calcula
  cuál es más óptimo para el algoritmo actual.
  Distancia total recorrida: 8297.28
  Tiempo de ejecución: 55 ms

### Comentarios finales

El problema me pareció muy interesante, me gustó mucho resolverlo. Es el primer
problema de optimización que resuelvo, y me gustó mucho la experiencia.

## Problema 2

### Impresiones del problema

En principio tengo las mismas impresiones que con el problema 1, la gran diferencia es
que este problema es mucho más largo. Esto hace que con la solución actual el tiempo de ejecución
sea muy alto, por lo que voy a intentar optimizarlo de alguna forma y va a ser un factor a tener
en cuenta además de la distancia total recorrida.

### Ideas para resolver el problema

Voy a empezar encarándolo de la misma forma que el problema 1, en caso de que el tiempo de ejecución
sea muy alto voy a volver al método de la primera entrega del problema 1 y en futuras entregas voy a
intentar optimizarlo más.

### Comentarios sobre el código

* Primer entrega:
  Como pensaba la primera solución del problema 1 toma demasiado tiempo de ejecución. Lo que hice fue emprolijar un poco
  el código y hacer que varios datos sean parametrizables, como por ejemplo la profundidad para calcular el recorrido.
  Para esta primera entrega elegí arbitrariamente la profundidad como 100. En próximas entregas voy a intentar optimizar
  el tiempo de ejecución para poder aumentar la profundidad y obtener un mejor resultado.
  Distancia total recorrida: 939034.02
  Tiempo de ejecución: 6 min 27 s
* Segunda entrega:
  En el algoritmo que planteé se calculan más de 340 millones de distancias, este número habría además que multiplicarlo
  por la profundidad que se elija. Esto hace que el tiempo de ejecución sea muy alto.
  ```math
  profundidad(n² - n) = #distancias que se calculan , donde n = cantidad de sucursales
  ```
  Quise agregar como optimización una caché de distancias que evite tener que calcular dos veces lo mismo, por ejemplo
  si ya calculé la distancia entre i y j no tendría sentido calcular la distancia entre j e i. Pero esto hizo que el
  tiempo de ejecución creciera muchísimo debido a que almacenar el total de distancias es demasiado. Igualmente,
  terminé notando que esto no tenía sentido, ya que al usar un slice de sucursales ya visitadas evito calcular dos veces
  lo mismo. Lo que si hice para optimizar fue agregar un criterio de corte más rápido para el cálculo de cada recorrido.
  Ahora le paso la mejor distancia a la función que se encarga de calcular el recorrido. De esta forma si el recorrido
  que se está calculando es peor simplemente se ignora. Usando esta optimización se mejora muy poco tiempo el de
  ejecución, pero noté mejoras en algunas pruebas. Para esta entrega aumenté la profundidad a 400.
  Distancia total recorrida: 931870.54
  Tiempo de ejecución: 29 min 11 s

### Comentarios finales

Sigue gustándome resolver este tipo de problemas, en este caso se agregó como dificultad el largo del
problema, lo que aumenta mucho el tiempo de ejecución. Es un gran desafío intentar optimizarlo.
Como conclusión, diría que dependiendo del problema y el hardware que se tenga, es importante hacer un
equilibrio entre la optimización y el tiempo de ejecución, por lo menos en el caso de la solución que
planteé.

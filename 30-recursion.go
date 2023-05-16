// Go acepta funciones recursivas. Una función es recursiva si se llama a sí misma y 
// alcanza una condición de parada.

// En el siguiente ejemplo, testcount() es una función que se llama a sí misma. 
// Utilizamos la variable x como datos, que se incrementa en 1 (x + 1) cada vez que hacemos una recursión. 
// La recursión termina cuando la variable x es igual a 11 (x == 11).

// package main
// import ("fmt")

// func testcount(x int) int {
//   if x == 18 {
//     return 0
//   }
//   fmt.Println(x)
//   return testcount(x + 1)
// }

// func main(){
//   testcount(1)
// } 

// La recursión es un concepto común en matemáticas y programación. 
// Esto tiene la ventaja de permitirte recorrer datos para llegar a un resultado.

// El desarrollador debe tener cuidado con las funciones recursivas, ya que es bastante fácil caer en la 
// escritura de una función que nunca termina, o que utiliza cantidades excesivas de 
// memoria o potencia de procesador. Sin embargo, cuando se escribe correctamente, 
// la recursión puede ser un enfoque muy eficiente y matemáticamente elegante para la programación.

// En el siguiente ejemplo, factorial_recursion() es una función que se llama a sí misma. 
// Utilizamos la variable x como datos, que se decrementa en 1 (-1) cada vez que hacemos una recursión. 
// La recursión termina cuando la condición no es mayor que 0 (es decir, cuando es 0).

package main

import "fmt"

func factorial_recursion(x float64) (y float64) {
  // La función factorial_recursion utiliza la recursión para calcular el factorial de un número.
  // Recibe un parámetro x de tipo float64 y devuelve un resultado de tipo float64.

  if x > 0 {
    // Si x es mayor que 0, la recursión continúa.
    // Multiplica x por el factorial de x-1 (llamada recursiva) y lo almacena en la variable y.
    y = x * factorial_recursion(x-1)
  } else {
    // Si x es igual o menor que 0, se considera que el factorial es 1.
    y = 1
  }

  return // Devuelve el valor almacenado en la variable y.
}

func main() {
  // En la función main, se llama a la función factorial_recursion y se imprime el resultado.
  fmt.Println(factorial_recursion(4))
}

// El resultado es 24. Esto se debe a que 4 * 3 * 2 * 1 = 24.
// En este caso, se llama a la función factorial_recursion con el argumento 4. 
// La función calcula el factorial de 4 utilizando recursión y devuelve el resultado. 
// Luego, el resultado se imprime en la salida estándar utilizando fmt.Println(). 
// En este ejemplo, el factorial de 4 es 24.

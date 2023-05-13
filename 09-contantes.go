// Si una variable debe tener un valor fijo que no se puede cambiar, puede usar el constpalabra clave.

// El palabra clave const declara la variable como "constante", 
// lo que significa que es inmutable y de solo lectura . 

// const CONSTNAME type = valu

package main
import ("fmt")

const PI = 3.14

func main() {

  fmt.Println(PI)

  // Hay 2 tipos de constantes:
  // Tipiadas
  const NUM int = 12
  fmt.Println(NUM)

  // No Tipiadas
  const NUM2 = 12
  fmt.Println(NUM2)

  // Cuando se declara una constante, no es posible cambiar el valor más tarde:
  // const NUM3 = 12
  // NUM3 = 13 // error

  // Se pueden agrupar varias constantes en un bloque para facilitar la lectura:
  const (
    NUM4 = 12
    NUM5 = 13
  )
  fmt.Println(NUM4, NUM5)

}

// Reglas

// Los nombres constantes siguen las mismas reglas que las variables

// Los nombres constantes generalmente se escriben en letras mayúsculas 
// (para una fácil identificación y diferenciación de las variables), BUENAS PRACTICAS

// Las constantes se pueden declarar tanto dentro como fuera de una función.









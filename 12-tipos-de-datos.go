// El tipo de datos es un concepto importante en la programación. 
// El tipo de datos especifica el tamaño y el tipo de los valores de las variables.

// Go tiene un tipo estático, lo que significa que una vez que se define un tipo de variable, 
// solo puede almacenar datos de ese tipo.

// Go tiene tres tipos de datos básicos:

//     bool : representa un valor booleano y es verdadero o falso
//     Numérico : representa tipos enteros, valores de coma flotante, y tipos complejos
//     string : representa un valor de cadena

package main
import ("fmt")

func main() {
  var a bool = true     // Boolean
  var b int = 5         // Integer
  var c float32 = 3.14  // Floating point number
  var d string = "Hi!"  // String

  fmt.Println("Boolean: ", a)
  fmt.Println("Integer: ", b)
  fmt.Println("Float:   ", c)
  fmt.Println("String:  ", d)
}

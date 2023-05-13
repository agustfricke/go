// Un archivo Go consta de las siguientes partes:

// Declaración de paquete
package main
// Importar paquetes
import "fmt"

// Declaración de funciones
func main() {
  // Declaraciones y expresiones
  fmt.Println("Hola Mundo")
}

// En Go, las declaraciones se separan al terminar una línea 
// (pulsando la tecla Intro) o al un punto y coma" ;".

// Presionar la tecla Enter agrega " ;" al final de la línea implícitamente 
// (no aparece en el código fuente).

// El corchete izquierdo { no puede venir al principio de una línea.

// El siguiente código no es válido:
// package main
// import ("fmt")

// func main()
// {
//   fmt.Println("Hello World!")
// }

// Puede escribir un código más compacto, como se muestra a continuación 
// (esto no se recomienda porque hace que el código más difícil de leer): 

// package main; import "fmt"; func main() { fmt.Println("Hello World!") }



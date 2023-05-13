// Go ofrece varios verbos de formato que se pueden usar con el Printf() función. 

// %v: Imprime el valor en el formato predeterminado.
// %#v: Imprime el valor en formato de sintaxis de Go.
// %T: Imprime el tipo del valor.
// %%: Imprime el signo de porcentaje (%).

package main
import ("fmt")

func main() {
  var i = 66.5
  var txt = "Hola Mundo!"

  fmt.Printf("%v\n", i)
  fmt.Printf("%#v\n", i)
  fmt.Printf("%v%%\n", i)
  fmt.Printf("%T\n", i)

  fmt.Printf("%v\n", txt)
  fmt.Printf("%#v\n", txt)
  fmt.Printf("%T\n", txt)
}


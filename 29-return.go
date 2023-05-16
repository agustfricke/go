// Si deseas que la función devuelva un valor, debes definir el tipo de datos del valor 
// de retorno (como int, string, etc.) 
// y utilizar la palabra clave "return" dentro de la función:

// code to be executed
// func FunctionName(param1 type, param2 type) type {
// return output
// } 

// Aquí, myFunction() recibe dos enteros (x e y) y devuelve su suma (x + y) como un entero (int):
// package main
// import ("fmt")

// func myFunction(x int, y int) int {
//   return x + y
// }

// func main() {
//   fmt.Println(myFunction(1, 2))
// }

// Aquí, nombramos al valor de retorno como resultado (de tipo int) y devolvemos el valor con un retorno 
// desnudo (lo que significa que usamos la declaración return sin especificar el nombre de la variable):
// package main
// import ("fmt")

// func myFunction(x int, y int) (result int) {
//   result = x + y
//   return
// }

// func main() {
//   fmt.Println(myFunction(1, 2))
// }

// Aquí, nombramos al valor de retorno como resultado (de tipo int) y devolvemos el valor con un retorno
// package main
// import ("fmt")

// func myFunction(x int, y int) (result int) {
//   result = x + y
//   return result
// }

// func main() {
//   fmt.Println(myFunction(1, 2))
// }

// package main
// import ("fmt")

// func myFunction(x int, y int) (result int) {
//   result = x + y
//   return result
// }

// func main() {
//   total := myFunction(1, 2)
//   fmt.Println(total)
// }

// Retorno múltiple
package main
import ("fmt")

func saludar(x int, y string) (result int, text string) {
  result = x + 1
  text = y + " Mundo"
  return
}

func main() {
  total, texto := saludar(1, "Hola")
  fmt.Println(total, texto)
  // O tambien algo asi 
  fmt.Println(saludar(2, "Hola"))
  // Si queres omitir el valor de retorno, puedes usar el guión bajo (_) como nombre de la variable:
  tot, _ := saludar(1, "Hola")
  fmt.Println(tot)
}










// Diferencias entre var y :=

// Var se puede usar dentro y fuera de las funciones
// Var se puede declarar y asignar valor por separado

// := solo se puede usar dentro de las funciones
// := no se puede declarar y asignar valor por separado

// Esto codigo funciona porque se esta declarando con var
package main
import ("fmt")

var a int
var b int = 2
var c = 3

func main() {
  a = 1
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}

// Este codigo no funciona porque se esta declarando con :=
// package main
// import ("fmt")

// a := 1

// func main() {
//   fmt.Println(a)
// }







// Podemos tener declaraciones if dentro de otras declaraciones if
// A esto se le llama if anidados

// if condition1 {
//      se ejecuta si condicion 1 es true
//   if condition2 {
//      se ejecuta si condicion1 y 2 son true
//   }
// } 

package main
import ("fmt")

func main() {
  num := 20
  if num >= 10 {
    fmt.Println("Num es mas grade que 10.")
    if num > 15 {
      fmt.Println("Num es mas grade que 15.")
     }
  } else {
    fmt.Println("Num es mas chico que 10.")
  }
}

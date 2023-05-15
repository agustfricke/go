// La declaracion else if nos permite evaluar multiples condiciones

// if condition 1 {
//    Se ejecuta si la condicion 1 es true
// } else if condition 2 {
//    Se ejecuta si condicion1 es false y la condicion2 es true 
// } else {
//    Se ejecuta si condicion 1 y 2 son false
// }

package main
import ("fmt")

func main() {
  a := 10
  b := 10
  if a < b {
    fmt.Println("a es menor que b.")
  } else if a > b {
    fmt.Println("a es mayor que b.")
  } else {
    fmt.Println("a y b son iguales.")
  }
}

// Si la condicion 1 y 2 son true se va a ejecutar la condicion 1
// package main
// import ("fmt")

// func main() {
//   x := 30
//   if x >= 10 {
//     fmt.Println("x es mas grade o igual que 10.")
//   } else if x > 20 {
//     fmt.Println("x es mas grade que 20.")
//   } else {
//     fmt.Println("x es menor que 10.")
//   }
// }

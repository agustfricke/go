// La declaracion else
// Utiliza el else para especificar un bloque de código Go que se ejecutará si NO se cumple una condición(false). 

// if condición {
  // Codigo va aqui si la condicion es true
// } else {
  // Code va aqui si la condicion es false
// }

package main
import ("fmt")

func main() {
  temperature := 12
  if (temperature > 20) {
    fmt.Println("Hace calor afuera")
  } else {
    fmt.Println("Hace frio afuera")
  }
}

// Asegurate que la sintaxis en la de arriba y no algo asi
// package main
// import ("fmt")

// func main() {
//   temperature := 14
//   if (temperature > 15) {
//     fmt.Println("It is warm out there.")
//   } // esto te va a dar un error
//   else {
//     fmt.Println("It is cold out there.")
//   }
// }

// El statement "switch"
// Utiliza el statement "switch" para seleccionar uno de varios bloques de código que se ejecutarán.

// Como funciona?
// La expresión se evalúa una vez.
// El valor de la expresión del "switch" se compara con los valores de cada caso.
// Si hay una coincidencia, se ejecuta el bloque de código asociado.
// La palabra clave "default" es opcional. Especifica un código para ejecutar si no hay ninguna coincidencia de caso.

// switch expression {
// case x:
//    // code block
// case y:
//    // code block
// case z:
// ...
// default:
//    // code block
// } 

// Ejemplo:
// A continuación utiliza un número de día de la semana para calcular el nombre del día de la semana:
// package main
// import ("fmt")

// func main() {
//   day := 4

//   switch day {
//   case 1:
//     fmt.Println("Lunes")
//   case 2:
//     fmt.Println("Martes")
//   case 3:
//     fmt.Println("Miercoles")
//   case 4:
//     fmt.Println("Jueves")
//   case 5:
//     fmt.Println("Viernes")
//   case 6:
//     fmt.Println("Sabado")
//   case 7:
//     fmt.Println("Domingo")
//   }
// }

// La palabra clave "default" especifica un código para ejecutar si no hay ninguna coincidencia de caso:
package main
import ("fmt")

func main() {
  dia := 8

  switch dia {
  case 1:
    fmt.Println("Lunes")
  case 2:
    fmt.Println("Martes")
  case 3:
    fmt.Println("Miercoles")
  case 4:
    fmt.Println("Jueves")
  case 5:
    fmt.Println("Viernes")
  case 6:
    fmt.Println("Sabado")
  case 7:
    fmt.Println("Domingo")
  default:
    fmt.Println("Numero de dia invalido")
  }
}

// Todos los valores de caso deben tener el mismo tipo que la expresión del "switch". 
// De lo contrario, el compilador generará un error.

// func main() {
//   x := 3
//   switch x {
//   case 1:
//     fmt.Println("Uno")
//   case "B":
//     fmt.Println("Dos")
// } ERROR





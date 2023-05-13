// En Go, existen diferentes tipos de variables, por ejemplo:

// int - almacena números enteros, como 123 o -123
// float32 - almacena números decimales, como 19.99 o -19.99
// string - almacena texto, como "Hola Mundo". Los valores de cadena están entre comillas dobles ("")
// bool - almacena valores con dos estados: verdadero o falso

// Hay 2 formas de declarar variables con go

// 1. Declarar variable con var
// var nombreVariable tipoVariable = valorVariable

// 2. Declarar variable con :=
// nombreVariable := valorVariable

package main
import "fmt"

func main() {
  // Algunos ejemplo con var
  var nombre string = "Agustin"
  var edad int = 22
  var altura float32 = 1.75
  var programador bool = true
  fmt.Println("Mi nombre es", nombre, "tengo", edad, "años, mido", altura, "y soy programador:", programador)

  // Algunos ejemplo con :=
  nombre2 := "Agustin"
  edad2 := 22
  altura2 := 1.75
  programador2 := true
  fmt.Println("Mi nombre es", nombre2, "tengo", edad2, "años, mido", altura2, "y soy programador:", programador2)

  // En este ejemplo, se están declarando tres variables: "estudiante", "estudiante2" y "x". 
  // La variable "estudiante" se declara con un tipo explícito de cadena ("string"), 
  // mientras que "estudiante2" y "x" se declaran sin especificar explícitamente su tipo. 

  // En Go, cuando una variable se declara sin especificar su tipo, 
  // el compilador infiere el tipo de datos basándose en el valor que se le asigna.

  var estudiante string = "Agustin" //type es string
  var estudiante2 = "Juan" //type es inferido
  x := 2 //type es inferido

  fmt.Println(estudiante) // Agustin
  fmt.Println(estudiante2) // Juan
  fmt.Println(x) // 2

  // // Declaracion de varaibales sin estado inicial
  var a string
  var b int
  var c bool

  fmt.Println(a) // ""
  fmt.Println(b) // 0
  fmt.Println(c) // false

  // // Declaracion y luego asignacion de variables
  var nombre3 string
  nombre3 = "Agustin"
  fmt.Println(nombre3) // Agustin
}








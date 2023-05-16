// Parámetros y argumentos

// La información se puede pasar a las funciones como un parámetro. 
// Los parámetros actúan como variables dentro de la función.

// Los parámetros y sus tipos se especifican después del nombre de la función, 
// dentro de paréntesis. Puedes agregar tantos parámetros como desees, solo sepáralos con una coma:

// func FunctionName(param1 type, param2 type, param3 type) {
  // code to be executed
// } 

package main
import ("fmt")

func intro(nombre string, edad int) {
  fmt.Println("Hola", nombre, "tienes", edad, "años")
}

func main() {
  intro("Juan", 30)
  intro("Maria", 25)
  intro("Pedro", 20)
}


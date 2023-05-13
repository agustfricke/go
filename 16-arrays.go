// Los arrays se utilizan para almacenar múltiples valores del mismo tipo en una sola variable

// Sintaxis con var

// var array_name = [length]datatype{values} // Aqui la longitud del array esta definido
// var array_name = [...]datatype{values} // Aqui no

// Sintaxis con :=

// array_name =: [length]datatype{values} // Aqui la longitud del array esta definido
// array_name =: [...]datatype{values} // Aqui no

// Ejemplo
package main
import ("fmt")

func main() {
  // Con longitud definida
  var arr1 = [5]int{1,2,3,4,5}
  arr2 := [5]int{6,7,8,9,10}

  fmt.Println(arr1)
  fmt.Println(arr2)

  // Sin la longitud definida
  var arr3 = [5]int{1,2,3,4,5}
  arr4 := [5]int{6,7,8,9,10}

  fmt.Println(arr3)
  fmt.Println(arr4)

  // Que pasa con los arrays de cadenas de carateres? Recuerda que el primer elemento comienza en 0
  var nombres = [3]string{"Juan", "Pedro", "Luis"}
  fmt.Println(nombres)

  // Acceder a los elementos del array
  paises := [3]string{"España", "Francia", "Italia"}

  fmt.Println(paises[0])
  fmt.Println(paises[2])

  // Cambiar elementos del array
  paises[2] = "Portugal"
  // Cambiamos el valor de Italia por Portugal
  fmt.Println(paises)

  // Declarar array sin valores?
  y:= [5]int{} // no inicializado
  z:= [5]int{1,2} // inicializado parcialmente

  fmt.Println(y)
  fmt.Println(z)

  // Inicializar un solo algunos elementos del array
  // El resto de los elementos se inicializan con 0, en string con ""
  // { indice: valor, indice: valor }
  saludo := [5]int{2:30,3:40}

  fmt.Println(saludo)

  // Como sabes que tan grande es un array?

  frutas := [3]string{"manzana", "pera", "naranja"}
  fmt.Println(len(frutas))
}



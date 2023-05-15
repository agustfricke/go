// Acceder a indice de un slice
// package main
// import ("fmt")

// func main() {
//   prices := []string{"Hola", "Mundo", "!"}

//   fmt.Println(prices[0])
//   fmt.Println(prices[2])
// }

// Cambiar valor de un indice de un slice
// package main
// import ("fmt")

// func main() {
//   nombres := []string{"Agustin", "Juan", "Pedro"}
//   fmt.Println(nombres)
//   // Cambiamos Pedro por Pablo
//   nombres[2] = "Pablo"
//   fmt.Println(nombres)
// }

// Agregar un elemento a un slice
// package main
// import "fmt"

// func main() {
//   nombres := []string{"Agustin", "Juan", "Pedro"}
//   fmt.Println(nombres)
//   // Agregamos un elemento al slice
//   nombres = append(nombres, "Pablo")
//   fmt.Println(nombres)
// }


// Juntar 2 slices
// V1 -- No se crea un 3 array
// package main
// import "fmt"

// func main() {
//   fruta := []string{"Manzana", "Banana", "Pera"}
//   fruta2 := []string{"Melon", "Fresa", "Naranja"}
//   fmt.Println(fruta)
//   fmt.Println(fruta2)
//   // Juntamos los slices
//   fruta = append(fruta, fruta2...)
//   fmt.Println(fruta)
// }

// v2 -- se crea un 3 array
// package main
// import ("fmt")

// func main() {
//   myslice1 := []int{1,2,3}
//   myslice2 := []int{4,5,6}
//   myslice3 := append(myslice1, myslice2...)

//   fmt.Printf("myslice3=%v\n", myslice3)
//   fmt.Printf("length=%d\n", len(myslice3))
//   fmt.Printf("capacity=%d\n", cap(myslice3))
// }

// Cambiar la longitud de un slice
// package main
// import ("fmt")

// func main() {
//   arr1 := [6]int{9, 10, 11, 12, 13, 14} // El array
//   myslice1 := arr1[1:5] // Slice el array
//   fmt.Printf("myslice1 = %v\n", myslice1)
//   fmt.Printf("length = %d\n", len(myslice1))
//   fmt.Printf("capacity = %d\n", cap(myslice1))

//   myslice1 = arr1[1:3] // Cambiar la longitud volviendo a rebanar el array
//   fmt.Printf("myslice1 = %v\n", myslice1)
//   fmt.Printf("length = %d\n", len(myslice1))
//   fmt.Printf("capacity = %d\n", cap(myslice1))

//   myslice1 = append(myslice1, 20, 21, 22, 23) // Cambiar la longitud agregando elementos
//   fmt.Printf("myslice1 = %v\n", myslice1)
//   fmt.Printf("length = %d\n", len(myslice1))
//   fmt.Printf("capacity = %d\n", cap(myslice1))
// }

// Eficiencia de la memoria

// Cuando se utilizan slices, Go carga todos los elementos subyacentes en la memoria.

// Si el array es grande y solo necesitas algunos elementos, 
// es mejor copiar esos elementos usando la función copy().

// La función copy() crea un nuevo array subyacente con solo los elementos 
// requeridos para el slice. Esto reducirá la memoria utilizada por el programa
// haciendo tu programa más eficiente.

package main
import ("fmt")

func main() {
  numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
  // Slice orginal
  fmt.Printf("numbers = %v\n", numbers)
  fmt.Printf("length = %d\n", len(numbers))
  fmt.Printf("capacity = %d\n", cap(numbers))

  // Crea una copia con solo los números necesarios
  neededNumbers := numbers[:len(numbers)-10] // Corta los últimos 10 elementos
  numbersCopy := make([]int, len(neededNumbers)) 
  copy(numbersCopy, neededNumbers)

  fmt.Printf("numbersCopy = %v\n", numbersCopy)
  fmt.Printf("length = %d\n", len(numbersCopy))
  fmt.Printf("capacity = %d\n", cap(numbersCopy))
}


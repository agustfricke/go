// Los slices son similares a los arrays, pero son más potentes y flexibles.

// Al igual que los arrays, los slices también se utilizan para almacenar varios valores del 
// mismo tipo en una sola variable.

// Sin embargo, a diferencia de los arrays, la longitud de un slice puede crecer y disminuir


// Hay 3 maneras de crear un slice:
// 1. Usando la sintaxis []datatype{values}
// 2. Usando un array existente
// 3. Usando la función make()

// 1. Usando la sintaxis []datatype{values}
// Declaramos un slice de enteros de longitud 3 y también la capacidad de 3.
// Hay dos funciones que se pueden usar para devolver la longitud y capacidad de un slice:
  
// len()función - devuelve la longitud del slice(el número de elementos del slice)
// cap()función - devuelve la capacidad del slice(la cantidad de elementos que el slice puede crecer) o reducir a)

// package main
// import ("fmt")

// func main() {
//   myslice1 := []int{}
//   fmt.Println(len(myslice1))
//   fmt.Println(cap(myslice1))
//   fmt.Println(myslice1)

//   myslice2 := []string{"Los", "Slices", "en", "go", "son", "poderosos"}
//   fmt.Println(len(myslice2))
//   fmt.Println(cap(myslice2))
//   fmt.Println(myslice2)
// }

// 2. Usando un array existente
// package main
// import ("fmt")

// func main() {
//   arr1 := [6]int{10, 11, 12, 13, 14,15}
//   myslice := arr1[2:4]

//   fmt.Printf("myslice = %v\n", myslice)
//   fmt.Printf("length = %d\n", len(myslice))
//   fmt.Printf("capacity = %d\n", cap(myslice))
// }

// 3. Usando la función make()
// package main
// import ("fmt")

// func main() {
//   myslice1 := make([]int, 5, 10)
//   fmt.Printf("myslice1 = %v\n", myslice1)
//   fmt.Printf("length = %d\n", len(myslice1))
//   fmt.Printf("capacity = %d\n", cap(myslice1))

//   myslice2 := make([]int, 5)
//   fmt.Printf("myslice2 = %v\n", myslice2)
//   fmt.Printf("length = %d\n", len(myslice2))
//   fmt.Printf("capacity = %d\n", cap(myslice2))
// }

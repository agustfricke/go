// Los mapas se utilizan para almacenar valores de datos en pares clave:valor.

// Cada elemento en un mapa es un par clave:valor.

// Un mapa es una colección desordenada y modificable que no permite duplicados.

// La longitud de un mapa es el número de sus elementos. Puedes encontrarla utilizando 
// la función len().

// El valor predeterminado de un mapa es nil.

// Los mapas contienen referencias a una tabla hash subyacente.

// Sintaxis de los mapas
// var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
// b := map[KeyType]ValueType{key1:value1, key2:value2,...}

// package main
// import ("fmt")

// func main() {
//   var personas = map[string]int{"John":25, "Jane":30, "Doe":40}
//   fmt.Println(personas)
// }

// Creando mapas usando la función make():

// package main
// import ("fmt")

// func main() {
//   var persona = make(map[string]string) // El map esta vacio
//   persona["name"] = "John"
//   persona["last_name"] = "Doe"
//   persona["age"] = "25"

//   city := make(map[string]int) // El map esta vacio
//   city["Oslo"] = 1
//   city["Madrid"] = 2
//   city["Berlin"] = 3

//   fmt.Printf("persona\t%v\n", persona)
//   fmt.Printf("city\t%v\n", city)
// }

// Este ejemplo muestra la diferencia entre declarar un mapa vacío 
// utilizando la función make() y hacerlo sin ella:
// package main
// import ("fmt")

// func main() {
//   var a = make(map[string]string)
//   var b map[string]string

//   fmt.Println(a == nil) // false
//   fmt.Println(b == nil) // true
// }


// Tipos de clave permitidos

// La clave de un mapa puede ser de cualquier tipo de dato para el cual esté definido el 
// operador de igualdad (==). Esto incluye:

// Booleanos
// Números
// Cadenas de texto
// Arreglos
// Punteros
// Estructuras
// Interfaces (siempre que el tipo dinámico admita la igualdad)
// Los tipos de clave no válidos son:

// Slices (rebanadas)
// Mapas
// Funciones
// Estos tipos no son válidos porque el operador de igualdad (==) no está definido para ellos.

// Tipos de valor permitidos

// Los valores de un mapa pueden ser de cualquier tipo de dato.

// Accediendo a elementos de un mapa
// value = map_name[key]

// package main
// import ("fmt")

// func main() {
//   var a = make(map[string]string)
//   a["brand"] = "Ford"
//   a["model"] = "Mustang"
//   a["year"] = "1964"

//   fmt.Printf(a["brand"])
// }

// Actualizando elementos de un mapa
// map_name[key] = value

// package main
// import ("fmt")

// func main() {
//   var a = make(map[string]string)
//   a["brand"] = "Ford"
//   a["model"] = "Mustang"
//   a["year"] = "1964"

//   fmt.Println(a)

//   a["year"] = "1970" // Updating an element
//   a["color"] = "red" // Adding an element

//   fmt.Println(a)
// }

// Eliminar elementos de un mapa con la función delete()
// delete(map_name, key)

// package main
// import ("fmt")

// func main() {
//   var a = make(map[string]string)
//   a["brand"] = "Ford"
//   a["model"] = "Mustang"
//   a["year"] = "1964"

//   fmt.Println(a)

//   delete(a,"year")

//   fmt.Println(a)
// }

// Verifica si una clave existe en un mapa
// val, ok := map_name[key]

// package main
// import ("fmt")

// func main() {
//   var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day":""}

//   val1, ok1 := a["brand"] // Checking for existing key and its value
//   val2, ok2 := a["color"] // Checking for non-existing key and its value
//   val3, ok3 := a["day"]   // Checking for existing key and its value
// Si solo deseas verificar la existencia de una determinada clave, puedes utilizar el identificador en blanco (_) en lugar de val.
//   _, ok4 := a["model"]    // Only checking for existing key and not its value

//   fmt.Println(val1, ok1)
//   fmt.Println(val2, ok2)
//   fmt.Println(val3, ok3)
//   fmt.Println(ok4)
// }

// Los mapas son referencias a tablas hash.

// Si dos variables de mapa se refieren a la misma tabla hash, al cambiar 
// el contenido de una variable afectará al contenido de la otra.

// package main
// import ("fmt")

// func main() {
//   var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
//   b := a

//   fmt.Println(a)
//   fmt.Println(b)

//   b["year"] = "1970"
//   fmt.Println("After change to b:")

//   fmt.Println(a)
//   fmt.Println(b)
// }

// Puedes usar la palabra clave range para iterar sobre mapas.
// Ejemplo con output desordenado:
// package main
// import ("fmt")

// func main() {
//   a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

//   for key, value := range a {
//     fmt.Printf("%v : %v, ", key, value)
//   }
// }

// Ejemplo con output ordenado:
package main

import (
	"fmt"
)

func main() {
	a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	b := []string{"one", "two", "three", "four"} // defining the order

	for _, v := range b { // loop with the defined order
		fmt.Printf("%v : %v, ", v, a[v])
	}
}












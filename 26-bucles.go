// El ciclo "for" recorre un bloque de código un número especificado de veces.
// El ciclo "for" es el único bucle disponible en Go.

// Los bucles son útiles si deseas ejecutar el mismo código una y otra vez, cada vez con un valor diferente.

// Cada ejecución de un bucle se llama una iteración.

// Sintaxis:
// for statement1; statement2; statement3 {
// code to be executed for each iteration
// }

// statement1: Inicializa el valor del contador del bucle.

// statement2: Se evalúa en cada iteración del bucle. 
// Si se evalúa como VERDADERO, el bucle continúa. Si se evalúa como FALSO, el bucle termina.

// statement3: Incrementa el valor del contador del bucle.

// Aqui un loop que cuenta hasta 10:
// package main
// import ("fmt")

// func main() {
//   for i:=0; i < 11; i++ {
//     fmt.Println(i)
//   }
// }

// Okey ahora contemos hasta 10 pero de 2 en 2:
// package main
// import ("fmt")

// func main() {
//   for i:=0; i < 11; i+=2 {
//     fmt.Println(i)
//   }
// }

// La declaración "continue" se utiliza para saltar una o más 
// iteraciones en el bucle. Luego continúa con la siguiente iteración en el bucle.
// Aqui un ejemplo donde saltamos el numero 3:

// package main
// import ("fmt")

// func main() {
//   for i := 0; i < 11; i++ {
//     if i == 3 {
//       continue
//     }
//     fmt.Println(i)
//   }
// }

// La declaración "break" se utiliza para romper/terminar la ejecución del bucle.
// Aqui un ejemplo donde rompemos el bucle cuando el contador es 3:
// package main
// import ("fmt")

// func main() {
//   for i := 0; i < 11; i++ {
//     if i == 3 {
//       break
//     }
//     fmt.Println(i)
//   }
// }



// Bucles anidados

// Es posible colocar un bucle dentro de otro bucle.

// Aquí, el "bucle interno" se ejecutará una vez por cada iteración del "bucle externo":
// package main
// import ("fmt")

// func main() {
//   adj := [2]string{"big", "tasty"}
//   fruits := [3]string{"apple", "orange", "banana"}
//   for i:=0; i < len(adj); i++ {
//     for j:=0; j < len(fruits); j++ {
//       fmt.Println(adj[i],fruits[j])
//     }
//   }
// }

// La palabra clave "range" se utiliza para iterar más fácilmente sobre un array, slice o mapa. 
// Devuelve tanto el índice como el valor.

// La palabra clave "range" se utiliza de la siguiente manera:

// for index, value := array|slice|map {
   // code to be executed for each iteration
// }
// package main
// import ("fmt")

// func main() {
//   fruits := [3]string{"apple", "orange", "banana"}
//   for idx, val := range fruits {
//      fmt.Printf("%v\t%v\n", idx, val)
//   }
// }
// Para mostrar únicamente el valor o el índice, puedes omitir la otra salida utilizando un guion bajo (_).

// package main
// import ("fmt")

// func main() {
//   fruits := [3]string{"apple", "orange", "banana"}
//   for _, val := range fruits {
//      fmt.Printf("%v\n", val)
//   }
// }

// O si quieremos mostrar solo el indice:
package main
import ("fmt")

func main() {
  frutas := [3]string{"manzana", "naranja", "banana"}
  for idx, _ := range frutas {
    fmt.Printf("%v\n", idx)
  }
}












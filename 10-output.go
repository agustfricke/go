// Go tiene tres funciones para generar texto: 
// Print, Println y Printf.

// El Print()La función imprime sus argumentos con su formato predeterminado.

package main
import ("fmt")

func main() {
  var pal1,pal2 string = "Hola","Mundo"

  fmt.Print(pal1)
  fmt.Print(pal2)

  // Si queremos imprimir los argumentos en líneas nuevas, necesitamos usar \n.
  fmt.Print(pal1, "\n")
  fmt.Print(pal2, "\n")

  // También es posible utilizar sólo uno Print()para imprimir múltiples variables. 
  fmt.Print(pal1, "\n",pal2)

  // Si queremos agregar un espacio entre argumentos de cadena, necesitamos usar " ": 
  fmt.Print(pal1, " ",pal2)

  // El Println()función es similar a Print()con la diferencia que 
  // se agrega un espacio en blanco entre los argumentos y se agrega una nueva línea al final: 

  fmt.Println(pal1, pal2)

  // El Printf()la función primero formatea su argumento en función del formato dado 
  // verbo y luego los imprime.

  // Aquí usaremos dos verbos de formato:

  // %vse utiliza para imprimir el valor de los argumentos
  // %Tse utiliza para imprimir el tipo de los argumentos

  fmt.Printf("pal1 has value: %v and type: %T\n", pal1, pal1)
  
}

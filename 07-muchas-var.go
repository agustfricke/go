package main
import ("fmt")

func main() {
  // En Go, es posible declarar múltiples variables en la misma línea. 
  var a, b, c, d int = 1, 3, 5, 7

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)

  // Si el type no se especifica la palabra clave, puede declarar diferentes tipos de 
  // variables en la misma línea:
  var num, nombre = 9, "Agustin"
  num2, nombre2 := 4, "Juan"

  fmt.Println(num, nombre)
  fmt.Println(num2, nombre2)

  // Las declaraciones de variables múltiples también se pueden 
  // agrupar en un bloque para mayor legibilidad: 
  var (
    nombre3 = "Agustin"
    edad = 22
    altura = 1.75
  ) 

  fmt.Println(nombre3, edad, altura)
}




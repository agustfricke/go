package main

// Adición: Suma dos valores x + y
// Resta: Resta un valor de otro x - y
// Multiplicación: Multiplica dos valores x * y
// División: Divide un valor entre otro x / y
// Módulo: Devuelve el resto de una división x % y
// Incremento: Aumenta el valor de una variable en 1 x++
// Decremento: Disminuye el valor de una variable en 1 x--

import "fmt"

func main() {
  // Operadores se asignación
  // Los operadores de asignación se utilizan para asignar valores a las variables. 
  var a int = 10
  fmt.Println(a) // 10
  // El operador de asignación de suma ( +=) añade un valor a una variable:
  a += 5
  fmt.Println(a) // 15

}

// Comparación
// == igual a x == y
// != no igual x != y
// mayor que x > y
// < menor que x < y
// = mayor o igual que x >= y
// <= menor o igual que x <= y


// Operadores lógicos
// && and && Retorna verdadero (true) si ambas declaraciones son verdaderas (true).
// || or Retorna verdadero (true) si una de las declaraciones es verdadera (true).
// ! not Revierte el resultado, retorna falso (false) si el resultado es verdadero (true).


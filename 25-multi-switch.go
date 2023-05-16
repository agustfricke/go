// El statement "switch" con múltiples casos
// Es posible tener múltiples valores para cada caso en el statement "switch":

// switch expression {
// case x,y:
//    // Evaluea x o y
// case v,w:
//    // Same para v o w
// case z:
// ...
// default:
//    // Si no hay coincidencia, se ejecuta
// } 

package main
import ("fmt")

func main() {
  dia := 2
  switch dia {
  case 1, 2, 3, 4, 5:
    fmt.Println("Dia de semana")
  case 6, 7:
    fmt.Println("Fin de semana")
  default:
    fmt.Println("Dia invalido")
  }
}



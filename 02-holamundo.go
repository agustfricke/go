// Escribamos nuestro primer programa con Go
// para eso debemos poner el siguiente comando en nuesta carpeta de trabajo
// go mod init ejemplo.com/holamundo
// Ahora crea un nuevo achivo llamado holamundo.go y empezemos a programar

package main
import "fmt"

func main() {
  fmt.Println("Hola Mundo")
}

// Ahora ejecutemos nuestro programa con el siguiente comando
// go run holamundo.go
// output: Hola Mundo

// Ahora vamos a compilar nuestro programa con el siguiente comando
// go build holamundo.go
// Para ejecutarlo en linux ./holamundo


// Las estructuras en Go

// Una estructura (abreviada como "struct" en inglés) se utiliza para crear una 
// colección de miembros de diferentes tipos de datos en una sola variable.

// Mientras que los arrays se utilizan para almacenar múltiples valores del mismo 
// tipo de dato en una sola variable, las estructuras se utilizan para almacenar múltiples 
// valores de diferentes tipos de datos en una sola variable.

// Una estructura puede ser útil para agrupar datos juntos y crear registros.

// Para declarar una estructura en Go, utiliza las palabras clave "type" y "struct":

// type struct_name struct {
//   member1 datatype;
//   member2 datatype;
//   member3 datatype;
// }

// Para acceder a cualquier miembro de una estructura, utiliza el operador punto (.) 
// entre el nombre de la variable de la estructura y el miembro de la estructura:
// package main

// import ("fmt")

// type Person struct {
//   name string;
//   age int;
//   dev bool;
// }

// func main() {

//   var person1 Person;
//   var person2 Person;

//   person1.name = "John";
//   person1.age = 25;
//   person1.dev = true;

//   person2.name = "Jane";
//   person2.age = 30;
//   person2.dev = false;

//   fmt.Println("Person 1 :", person1.name, person1.age, person1.dev);
//   fmt.Println("Person 2 :", person2.name, person2.age, person2.dev);

// }

// También puedes pasar una estructura como argumento de una función, de la siguiente manera:

package main
import ("fmt")

type Person struct {
  name string;
  age int;
  dev bool;
}

func printPerson(person Person) {
  fmt.Println("Name:", person.name);
  fmt.Println("Age:", person.age);
  fmt.Println("Developer:", person.dev);
}

func main() {
    var person1 Person;
    var person2 Person;
  
    person1.name = "John";
    person1.age = 25;
    person1.dev = true;
  
    person2.name = "Jane";
    person2.age = 30;
    person2.dev = false;
  
    printPerson(person1);
    printPerson(person2);
}











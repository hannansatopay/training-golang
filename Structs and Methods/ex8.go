package main
import "fmt"

func main() {

  var person struct {
    name, surname string
  }

  person.name, person.surname = "Barack", "Obama"
  anotherPerson := struct {   // anonymous struct
    name, surname string
  }{"Barack", "Obama"}
  fmt.Println(person, anotherPerson)
}
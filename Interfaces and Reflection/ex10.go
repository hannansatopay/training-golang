package main
import "fmt"

var i = 5
var str = "ABC"

type Person struct {
	name string
	age  int
}

type Any interface{}	// empty interface

func main() {
	var val Any
	val = 5	// assigning integer to empty interface
	fmt.Printf("val has the value: %v\n", val)
	val = str // assigning string to empty interface
	fmt.Printf("val has the value: %v\n", val)
	pers1 := new(Person)
	pers1.name = "Rob Pike"
	pers1.age = 55
	val = pers1 // assigning *Person type variable to empty interface
	fmt.Printf("val has the value: %v\n", val)
	switch t := val.(type) {	// cases defined on type of val
	case int:	// if val is int
		fmt.Printf("Type int %T\n", t)
	case string: // if val is string
		fmt.Printf("Type string %T\n", t)
	case bool: // if val is bool
		fmt.Printf("Type boolean %T\n", t)
	case *Person:	// if val is *Person
		fmt.Printf("Type pointer to Person %T\n", *t)
	default:	// None of the above types
		fmt.Printf("Unexpected type %T", t)
	}
}
package main
import "fmt"

type specialString string

var whatIsThis specialString = "hello"

func TypeSwitch() {
	testFunc := func(any interface{}) {	// lambda function in combination with empty interface
		switch v := any.(type) {
		case bool: // if v is bool
			fmt.Printf("any %v is a bool type", v)
		case int: // if v is int
			fmt.Printf("any %v is an int type", v)
		case float32: // if v is float32
			fmt.Printf("any %v is a float32 type", v)
		case string: // if v is string
			fmt.Printf("any %v is a string type", v)
		case specialString: // if v is specialString
			fmt.Printf("any %v is a special String!", v)
		default: // none of types satisfied
			fmt.Println("unknown type!")
		}
	}
	testFunc(whatIsThis)
}

func main() {
	TypeSwitch()
}
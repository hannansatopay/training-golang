package main
import (
	"fmt"
	"reflect"
	"strings"
)

func mapToStruct(m map[string]interface{}) interface{} {
	var structFields []reflect.StructField

	for k, v := range m {
		sf := reflect.StructField{
			Name: strings.Title(k),
			Type: reflect.TypeOf(v),
		}
		structFields = append(structFields, sf)
	}

	// Creates the struct type
	structType := reflect.StructOf(structFields)

	// Creates a new struct
	structObject := reflect.New(structType)

	return structObject.Interface()
}


func main() {

	m := make(map[string]interface{})

	m["name"] = "Barack"
	m["surname"] = "Obama"
	m["age"] = 57

	s := mapToStruct(m)
	fmt.Printf("%t %[1]v\n", s)

	sr := reflect.ValueOf(s)
	sr.Elem().FieldByName("Name").SetString("Donald")
	sr.Elem().FieldByName("Surname").SetString("Trump")
	sr.Elem().FieldByName("Age").SetInt(72)
	fmt.Println(s)
}
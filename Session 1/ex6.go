package main
import (
"fmt"
"reflect"
)

type T struct {
  a int "This is a tag"
  b int `A raw string tag`
  c int `key1:"value1" key2:"value2"`
}

func main() {
  t := T{}
  fmt.Println(reflect.TypeOf(t).Field(0).Tag)
  if field, ok := reflect.TypeOf(t).FieldByName("b"); ok {
    fmt.Println(field.Tag)
  }
  if field, ok := reflect.TypeOf(t).FieldByName("c"); ok {
    fmt.Println(field.Tag)
    fmt.Println(field.Tag.Get("key1"))
  }
  if field, ok := reflect.TypeOf(t).FieldByName("d"); ok {
    fmt.Println(field.Tag)
  } else {
    fmt.Println("Field not found")
  }
}
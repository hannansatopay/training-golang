package main
import "fmt"

type S struct {
  a int
}

type SType S // New type
type SAlias = S // Alias
type IntType int // New type
type IntAlias = int // Alias

func (recv S) print() { // function for type defined on type S
  fmt.Printf("%t: %[1]v\n", recv)
}
func (recv SType) print() { // function for type defined on the basis of S
  fmt.Printf("%t: %[1]v\n", recv)
}

// func (recv SAlias) print() { // <-- error: S.print redeclared in this block previous declaration at ./struct_method.go:15:6
// fmt.Printf("%t: %[1]v\n", recv)
// }

func (recv IntType) print() { // function for type defined on type on the basis of int
  fmt.Printf("%t: %[1]v\n", recv)
}

// func (recv IntAlias) print() { // <-- error: cannot define new methods on non-local type int
// fmt.Printf("%t: %[1]v\n", recv)
// }

func main(){
  a := S{10}
  a.print() // calling function from line 13
  b := SType{20}
  b.print() // calling function from line 16
  c := SAlias{30}
  c.print() // calling function from line 13
  d := IntType(40)
  d.print() // calling function from line 24
  // e := IntAlias(50) <-- error: e.print undefined (type int has no field or method print)
  // e.print()
}
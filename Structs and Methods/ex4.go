package main
import (
  "fmt"
  "unsafe"  // to use function Sizeof()
)

type T1 struct {  
  a, b int64
}

type T2 struct {
  t1 *T1  // pointer to T1
}

type T3 struct {
  t1 T1   // value type of T1
}

func main() {
  fmt.Println("Size of T1:", unsafe.Sizeof(T1{}))       // T1 value type
  fmt.Println("Size of T2:", unsafe.Sizeof(T2{&T1{}}))  // T2 containing pointer to T1
  fmt.Println("Size of T3:", unsafe.Sizeof(T3{}))       // Value of T3
}
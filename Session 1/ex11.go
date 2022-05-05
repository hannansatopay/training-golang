package main
import "fmt"

type TwoInts struct {
  a int
  b int
}

func main() {
  two1 := new(TwoInts)
  two1.a = 12
  two1.b = 10
  fmt.Printf("The sum is: %d\n", two1.AddThem()) // calling method
  fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20)) // calling method
  two2 := TwoInts{3, 4}
  fmt.Printf("The sum is: %d\n", two2.AddThem()) // calling method
}

func (tn *TwoInts) AddThem() int { // can be called by pointer to TwoInt type var.
  return tn.a + tn.b
}

func (tn *TwoInts) AddToParam(param int) int { // can be called by pointer to TwoInt type var.
  return tn.a + tn.b + param
}
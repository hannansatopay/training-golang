package main
import (
"fmt"
"strconv"
)

type TwoInts struct {
  a int
  b int
}

func main() {
  two1 := new(TwoInts)
  two1.a = 12
  two1.b = 10
  fmt.Printf("two1 is: %v\n", two1) // output: two1 is: (12 / 10)
  fmt.Println("two1 is:", two1) // output: two1 is: (12 / 10)
  fmt.Printf("two1 is: %T\n", two1) // output: two1 is: *main.TwoInts
  fmt.Printf("two1 is: %#v\n", two1) // output: &main.TwoInts{a:12, b:10}
  }

  func (tn *TwoInts) String() string {
    return "(" + strconv.Itoa(tn.a) + " / " + strconv.Itoa(tn.b) + ")"
}
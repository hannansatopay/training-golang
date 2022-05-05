package main
import "fmt"

type T struct {
  a int
}

func (t T) print(message string) {
  fmt.Println(message, t.a)
}

func (T) hello(message string) {
  fmt.Println("Hello!", message)
}

func callMethod(t T, method func(T, string)) {
  method(t, "A message")
}

func main() {
  t1 := T{10}
  t2 := T{20}
  var f func(T, string) = T.print
  callMethod(t1, f)
  callMethod(t2, f)
  callMethod(t1, T.hello)
}
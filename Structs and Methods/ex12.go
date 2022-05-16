package main
import (
"fmt"
"time"
)

type myTime struct {
  time.Time //anonymous field
}

func (t myTime) first3Chars() string {
  return t.String()[0:3]
}

func main() {
  m := myTime{time.Now()}
  //calling existing String method on anonymous Time field
  fmt.Println("Full time now:", m.String())
  fmt.Println("First 3 chars:", m.first3Chars()) //calling myTime.first3Chars
}
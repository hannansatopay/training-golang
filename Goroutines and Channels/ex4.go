package main
import (
  "fmt"
  "time"
)

func main() {
  ch1 := make(chan int)
  go pump(ch1) // pump hangs
  fmt.Println(<-ch1) // prints only 0
  time.Sleep(1e9)
}

func pump(ch chan int) {
  for i:= 0; ; i++ {
  ch <- i
  }
}
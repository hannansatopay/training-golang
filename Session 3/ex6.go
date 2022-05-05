package main
import (
  "fmt"
  "time"
)

func main() {
  stream := pump()
  go suck(stream)
  // the above 2 lines can be shortened to: go suck( pump() )
  time.Sleep(1e9)
}

func pump() chan int {
  ch := make(chan int)
  go func() {
    for i := 0; ; i++ {
      ch <- i
    }
  }()
  return ch
}

func suck(ch chan int) {
  for {
    fmt.Println(<-ch)
  }
}
package main
import (
  "fmt"
  "time"
)

func main() {
  ch := make(chan string)
  go sendData(ch) // calling goroutine to send the data
  go getData(ch)  // calling goroutine to receive the data
  time.Sleep(1e9)
}

func sendData(ch chan string) { // sending data to ch channel
  ch <- "Washington"
  ch <- "Tripoli"
  ch <- "London"
  ch <- "Beijing"
  ch <- "Tokyo"
}

func getData(ch chan string) {
  var input string
  for {
    input = <-ch // receiving data sent to ch channel
    fmt.Printf("%s ", input)
  }
  close(ch) // closed the channel
}
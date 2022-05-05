package main
import "fmt"

type Camera struct { }

func (c *Camera) TakeAPicture() string {  // method of Camera
  return "Click"
}

type Phone struct { }

func (p *Phone ) Call() string {  // method of Phone
  return "Ring Ring"
}

// multiple inheritance
type SmartPhone struct {    // can use methods of both structs
  Camera
  Phone
}

func main() {
  cp := new(SmartPhone)
  fmt.Println("Our new SmartPhone exhibits multiple behaviors ...")
  fmt.Println("It exhibits behavior of a Camera: ", cp.TakeAPicture())
  fmt.Println("It works like a Phone too: ", cp.Call())
}
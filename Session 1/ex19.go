package main
import (
"fmt"
"math"
)

type Point struct {
  x, y float64
}

func (p *Point) Abs() float64 {
  return math.Sqrt(p.x*p.x + p.y*p.y)
}

type NamedPoint struct {
  Point   // anonymous field of Point type
  name string
}

func main() {
  n := &NamedPoint{Point{3, 4}, "Pythagoras"} // making pointer type variable
  fmt.Println(n.Abs()) // prints 5
}
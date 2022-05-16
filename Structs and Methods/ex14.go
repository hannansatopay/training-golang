package main
import "fmt"

type Rectangle struct { // struct of type Rectangle
    length, width int
}

func (r *Rectangle) Area() int {    // method calculating area of rectangle
    return r.length * r.width
}

func (r *Rectangle) Perimeter() int { // method calculating perimeter of rectangle
    return 2* (r.length + r.width)
}

func main() {
    r1 := Rectangle{4, 3}
    fmt.Println("Rectangle is: ", r1)
    fmt.Println("Rectangle area is: ", r1.Area()) // calling method of area
    fmt.Println("Rectangle perimeter is: ", r1.Perimeter()) // calling method of perimeter
}
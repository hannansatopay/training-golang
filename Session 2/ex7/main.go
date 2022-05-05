package main
import (
	"fmt"
	// "sort"      // this uses the Go sort package, then replace mysort. with sort. in the code below
	"./mysort" // this uses our own sort package (a subset of the Go sort package)
)

// sorting of slice of integers
func ints() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	a := mysort.IntSlice(data) //conversion to type IntSlice
	mysort.Sort(a)
	if !mysort.IsSorted(a) {
		panic("fail")
	}
	fmt.Printf("Numbers: %v\n", a)
}

// sorting of slice of strings
func strings() {
	data := []string{"Monday", "Friday", "Tuesday", "Wednesday", "Sunday", "Thursday", "", "Saturday"}
	a := mysort.StringSlice(data)
	mysort.Sort(a)
	if !mysort.IsSorted(a) {
		panic("fail")
	}
	fmt.Printf("Alphabets: %v\n", a)
}

// a type which describes a day of the week
type day struct {
	num       int
	shortName string
	longName  string
}

type dayArray struct {
	data []*day
}

func (p *dayArray) Len() int           { return len(p.data) }
func (p *dayArray) Less(i, j int) bool { return p.data[i].num < p.data[j].num }
func (p *dayArray) Swap(i, j int)      { p.data[i], p.data[j] = p.data[j], p.data[i] }

// sorting of custom type day
func days() {
	Sunday := day{6, "SUN", "Sunday"}
	Monday := day{0, "MON", "Monday"}
	Tuesday := day{1, "TUE", "Tuesday"}
	Wednesday := day{2, "WED", "Wednesday"}
	Thursday := day{3, "THU", "Thursday"}
	Friday := day{4, "FRI", "Friday"}
	Saturday := day{5, "SAT", "Saturday"}
	data := []*day{&Tuesday, &Thursday, &Wednesday, &Sunday, &Monday, &Friday, &Saturday}
	a := dayArray{data}
	mysort.Sort(&a)
	if !mysort.IsSorted(&a) {
		panic("fail")
	}
	for _, d := range data {
		fmt.Printf("%s ", d)
	}
	fmt.Printf("\n")
}

func main() {
	ints()
	strings()
	days()
}
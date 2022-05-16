# Regular Mistakes and Suggestions

## Common pitfalls

-   Never use  `var p*a`  to not confuse pointer declaration and multiplication.
-   Never change the counter-variable within the for-loop itself.
-   Never use the value in a for-range loop to change the value itself.
-   Never use  `goto`  with a preceding label.
-   Never forget the  _braces_  () after a function name to call the function specifically when calling a method on a receiver or invoking a lambda function as a Goroutine.
-   Never use  `new()`  with maps, always use  `make()`.
-   When coding a  `String()`  method for a type, don’t use  `fmt.Print`  or alike in the code, use the  `fmt.Sprint`  family of methods.
-   Never forget to use  `Flush()`  when terminating buffered writing.
-   Never ignore errors because ignoring them can lead to program crashes.
-   Do not use global variables or shared memory; they make your code unsafe for running concurrently.
-   When using  _JSON_, make sure that the data fields you require are exported in the data structure.
-   Use  `println`  only for debugging purposes.

## Best practices

In contrast, use the following:

-   Initialize a slice of maps the right way.
-   Always use the “comma, ok” (or checked) form for type assertions.
-   Make and initialize your types with a factory.
-   Use a pointer as receiver for a method on a struct only when the method modifies the structure, otherwise use a value.
-   Always use the package  `html/template`  to implement data-driven templates for generating HTML output safe against code injection.

## Hiding a variable by misusing short declaration

In the following code snippet:

```go
var remember bool = false
if something {
	remember := true // Wrong
}
// use remember
```

The variable  `remember`  will never become true outside of the if-body. Inside the if-body, a new remember variable that hides the outer remember is declared because of  `:=`, and there it will be true. However, after the  _closing }_  of  `if`, the variable  `remember`  regains its outer value  _false_. So, write it as:

```go
if something {
	remember = true
}
```

This can also occur with a for-loop, and can be particularly subtle in functions with  _named return variables_, as the following snippet shows:

```go
func shadow() (err error) {
  x, err := check1() // x is created; err is assigned to
  if err != nil {
    return // err correctly returned
  }
  if y, err := check2(x); err != nil { // y and inner err are created
    return // inner err shadows outer err so err is wrongly returned!
  } else {
    fmt.Println(y)
  }
  return
}
```

## Misusing strings

When you need to do a lot of manipulations on a string, think about the fact that strings in Go (like in Java and C#) are immutable. String concatenations of the kind  `a += b`  are inefficient, mainly when performed inside a loop. They cause many reallocations and the copying of memory. Instead, one should use a  `bytes.Buffer`  to accumulate string content, like in the following snippet:

```go
var b bytes.Buffer
...
for condition {
  b.WriteString(str) // appends string str to the buffer
}
return b.String()
```

> **Remark**: Due to compiler-optimizations and the size of the strings, using a buffer only starts to become more efficient when the number of concatenations is  **> 15**.

## Using  `defer`  for closing a file in the wrong scope

Suppose you are processing a range of files in a  _for_-loop, and you want to make sure the files are closed after processing by using  `defer`, like this:

```go
for _, file := range files {
  if f, err = os.Open(file); err != nil {
    return
  }
  // This is wrong. The file is not closed when this loop iteration ends.
  defer f.Close()
  // perform operations on f:
  f.Process(data)
}
```

But, at the end of the for-loop,  `defer`  is not executed. Therefore, the files are not closed! Garbage collection will probably close them for you, but it can yield errors. Better do it like this:

```go
for _, file := range files {
  if f, err = os.Open(file); err != nil {
    return
  }
  // perform operations on f:
  f.Process(data)
  // close f:
  f.Close()
}
```

_The keyword  `defer`  is only executed at the return of a function, not at the end of a loop or some other limited scope_.

## Confusing  `new()`  and  `make()`

-   For slices, maps and channels: use  `make`
-   For arrays, structs and all value types: use  `new`

## No need to pass a pointer to a slice to a function

A slice is a pointer to an underlying array. Passing a slice as a parameter to a function is probably what you always want, which means namely passing a pointer to a variable to be able to change it, and not passing a copy of the data. So, you want to do this:

```go
func findBiggest( listOfNumbers []int ) int {}
```

not this:

```go
func findBiggest( listOfNumbers *[]int ) int {}
```

_Do not dereference a slice when used as a parameter!_

## Using pointers to interface types

Look at the following program which  _can’t_  be compiled:

```go
package main
import (
"fmt"
)

type nexter interface {
  next() byte
}

func nextFew1(n nexter, num int) []byte {
  var b []byte
  for i:=0; i < num; i++ {
    b[i] = n.next()
  }
  return b
}

func nextFew2(n *nexter, num int) []byte {
  var b []byte

  for i:=0; i < num; i++ {
    b[i] = n.next()
    // compile error:

    // n.next undefined (type *nexter is pointer to interface, not interface)
  }
  return b
}

func main() {
  fmt.Println("Hello World!")
}
```

In the code above,  `nexter`  is an interface with the method  `next()`, which reads the next byte.  `nextFew1()`  has this interface type as a parameter and reads the next  `num`  bytes, returning them as a slice. This is ok. However,  `nextFew2()`  uses a pointer to the interface type as a parameter. When using the  `next()`  function, we get a clear compiler error:  `n.next undefined (type *nexter is pointer to interface, not interface)`.

_So never use a pointer to an interface type; this is already a pointer!_

## Misusing pointers with value types

Passing a value as a parameter in a function or as a receiver to a method may seem a misuse of memory because a value is always copied. But, on the other hand, values are allocated on the stack, which is quick and relatively cheap. If you pass a pointer to the value instead of the Go compiler, in most cases, we will see this as the making of an object, and we will move this object to the heap, causing an additional memory allocation. Therefore, nothing was gained in using a pointer instead of the value.

## Misusing goroutines and channels

For didactic reasons and for gaining insight into their working, a lot of the examples in applied goroutines and channels in very simple algorithms, like as a generator or iterator. In practice, often, you don’t need the concurrency, or you don’t need the overhead of the goroutines with channels; passing parameters using the stack is, in many cases, far more efficient. Moreover, it is likely to leak memory if you break, return or panic your way out of the loop because the goroutine blocks in the middle of doing something. In real code, it is often better to just write a simple procedural loop. Use goroutines and channels only where concurrency is important!

## Using closures with goroutines

```go
package main

import (

"fmt"

"time"

)

  

var values = [5]int{10, 11, 12, 13, 14}

  

func main() {

// version A:

for ix := range values { // ix is the index

func() {

fmt.Print(ix, " ")

}() // call closure, prints each index

}

fmt.Println()

// version B: same as A, but call closure as a goroutine

for ix := range values {

go  func() {

fmt.Print(ix, " ")

}()

}

  

fmt.Println()

time.Sleep(5e9)

// version C: the right way

for ix := range values {

go  func(ix int) {

fmt.Print(ix, " ")

}(ix)

}

fmt.Println()

time.Sleep(5e9)

// version D: print out the values:

for ix := range values {

val := values[ix]

go  func() {

fmt.Print(val, " ")

}()

}

time.Sleep(1e9)

}
```

**Version A**  calls a closure five times, which prints the value of the index.  **Version B**  does the same but invokes each closure as a goroutine, with an argument that this would be faster because the closures execute in parallel. If we leave enough time for all goroutines to run, the output of version B is  **4 4 4 4 4**. Why is this? The  `ix`  variable in the above loop B is a single variable that takes on the index of each array element. Because the closures are all only bound to that one variable, there is a very good chance that, when you run this code, you will see the last index (4) printed for every iteration, instead of each index in the sequence. This is because the goroutines will probably not begin executing until after the loop, when  `ix`  has the value 4.

The right way to code that loop is  **version C**: invoke each closure with  `ix`  as a parameter. Then,  `ix`  is evaluated at each iteration and placed on the stack for the goroutine, so each index is available to the goroutine when it is eventually executed. Note that the output depends on when each of the goroutines starts.

In  **version D**, we print out the values of the array. Why does this work and version B do not? Because variables declared within the body of a loop (as  `val`  here) are not shared between iterations, and thus can be used separately in a closure.

## Bad error handling

### Don’t use booleans

Making a  _boolean_  variable whose value is a test on the error-condition, like in the following, is superfluous:

```go
var good bool
// test for an error, good becomes true or false
if !good {
  return errors.New("things aren't good")
}
```

Instead, test on the error immediately:

```go
_, err1 := api.Func1()
if err1 != nil { ... }
```

### Don’t clutter code with error-checking

Avoid writing code like this:

```go
... err1 := api.Func1()
if err1 != nil {
  fmt.Println("err: " + err.Error())
  return
}
err2 := api.Func2()
if err2 != nil {
  ...
  return
}
```

First, include the call to the functions in an initialization statement of the  _if’s_. Even then, the errors are reported (by printing them) with if-statements scattered throughout the code. With this pattern, it is hard to tell what is normal program logic and what is error checking/reporting. Also, notice that most of the code is dedicated to error conditions at any point in the code. A good solution is to wrap your error conditions in a closure wherever possible, like in the following example:

```go
func httpRequestHandler(w http.ResponseWriter, req *http.Request) {
  err := func () error {
    if req.Method != "GET" {
      return errors.New("expected GET")
    }
    if input := parseInput(req); input != "command" {
      return errors.New("malformed command")
    }
    // other error conditions can be tested here
   } ()
   if err != nil {
     w.WriteHeader(400)
     io.WriteString(w, err)
     return
   }
  doSomething() ...
  ...
```

This approach clearly separates the error checking, error reporting, and normal program logic.

## for  `range`  in arrays

If you make a for  `range`  over an array with 1 value after the for, you get the indices back, not the values of the array. In:

```go
for n := range array { ... }
```

The variable  `n`  gets the values  **0, 1, … , len(array)-1**. Moreover, the compiler won’t warn you if the operations in the loop are compatible with  _ints_.

# The comma, ok Pattern

While studying the Go-language, we encountered several times the so-called  _comma, ok_  idiom where an expression returns two values: the first of which is a  _value_  or  _nil_, and the second is  _true/false_  or an error. An if-condition with initialization and then testing on the second-value leads to succinct and elegant code. This is a significant pattern in idiomatic Go-code. Here are all cases summarized:

## Testing for errors on function return

```go
var value Type_value
var err error
if value, err = pack1.Func1(param1); err != nil {
  fmt.Printf("Error %s in pack1.Func1 with parameter %v", err.Error(), param1)
  return err
}
// no error in Func1:
Process(value)```

Other places where it is used:

```go
os.Open(file), strconv.Atoi(str)
```

The following code will not compile because the scope of the file variable is the if-block only:

```go
if file, err := os.Open("input.dat"); err != nil {
  fmt.Printf("An error occurred on opening the inputfile\n" +
  "Does the file exist?\n" +
  "Have you got access to it?\n")
  return // exit the function on error
}
defer file.Close()
reader := bufio.NewReader(file)```
```
In this case, declare the variables upfront:

```go
var file *os.File
var err error
```

and then  `file, err := os.Open("input.dat")`  can be replaced by:  `file, err = os.Open("input.dat")`.

The function in which this code occurs returns the error to the caller, giving it the value nil when the normal processing was successful and so has the signature:

```go
func SomeFunc() error {
  ...
  if value, err = pack1.Func1(param1); err != nil {
    ...
    return err
  }
  ...
  return nil
}
```
The same pattern is used when recovering from a  _panic_  with  `defer`. A good pattern for clean error checking is using closures.

## Testing if a key-value item exists in a map

Does this means that  `map1`  has a value for  `key1`? Look at the following snippet:

```go
if value, isPresent = map1[key1]; isPresent {
  Process(value)
}
// key1 is not present
...
```
## Testing if an interface variable is of certain type

```go
if value, ok := varI.(T); ok {
  Process(value)
}
// varI is not of type T
```

## Testing if a channel is closed

```go
for input := range ch {
  Process(input)
}
```

or:

```go
var input Type_of_input
var err error
for {
  if input, open = <-ch; !open {
    break // channel is closed
  }
  Process(input)
}
```

# The defer Pattern
Using  `defer`  ensures that all resources are properly closed or given back to the  _pool_  when the resources are not needed anymore. Secondly, it is paramount in Go to recover from panicking.

## Closing a file stream

```go
// open a file f
defer f.Close()
```

## Unlocking a locked resource (a mutex)

```go
mu.Lock()
defer mu.Unlock()
```

## Closing a channel (if necessary)

```go
ch := make(chan float64)
defer close(ch)
```

or with 2 channels:

```go
answerα, answerβ := make(chan int), make(chan int)
defer func() { close(answerα); close(answerβ) }()
```

## Recovering from a panic

```go
defer func() {
  if err := recover(); err != nil {
    log.Printf("run time panic: %v", err)
}()
```

## Stopping a ticker

```go
tick1 := time.NewTicker(updateInterval)
defer tick1.Stop()
```

## Release of a process

```go
p, err := os.StartProcess(..., ..., ...)
defer p.Release()
```

## Stopping CPU profiling and flushing the info[#]

```go
pprof.StartCPUProfile(f)
defer pprof.StopCPUProfile()
```

## A goroutine signaling a WaitGroup

```go
func HeavyFunction1(wg *sync.WaitGroup) {
defer wg.Done()
// Do a lot of stuff
}
```

It can also be used when not forgetting to print a footer in a report.

# Visibility and Operator Patterns

## The operator pattern and interfaces

An operator is a  _unary_  or  _binary_  function which returns a new object and does not modify its parameters, like + and *. In C++, special infix operators (+, -, *, and so on) can be overloaded to support math-like syntax, but apart from a few special cases Go does not support operator overloading. To overcome this limitation, operators must be simulated with functions. Since Go supports a procedural as well as an object-oriented paradigm, there are two options:

### Implement the operators as functions

The operator is implemented as a package-level function to operate on one or two parameters and return a new object, implemented in the package dedicated to the objects on which they operate. For example, if we implement a matrix manipulation in a package  `matrix`, this would contain the addition of matrices  `Add()`  and multiplication  `Mult()`, which result in a matrix. These would be called on the package name itself so that we could make expressions of the form:

```go
m := matrix.Add(m1, matrix.Mult(m2, m3))
```

If we would like to differentiate between the different kinds of matrices (sparse and dense) in these operations, because there is no function overloading, we would have to give them different names, as in:

```go
func addSparseToDense (a *sparseMatrix, b *denseMatrix) *denseMatrix
func addDenseToDense (a *denseMatrix, b *denseMatrix) *denseMatrix
func addSparseToSparse (a *sparseMatrix, b *sparseMatrix) *sparseMatrix
```

This is not very elegant, and the best we can do is hide these as private functions and expose a single public function  `Add()`. This can operate on any combination of supported parameters by type-testing them in a nested type switch:

```go
func Add(a Matrix, b Matrix) Matrix {
  switch a.(type) {
  case sparseMatrix:
    switch b.(type) {
    case sparseMatrix:
      return addSparseToSparse(a.(sparseMatrix), b.(sparseMatrix))
    case denseMatrix:
      return addSparseToDense(a.(sparseMatrix), b.(denseMatrix))
    ...
  default:
    // unsupported arguments
    ...
  }
}
```

However, the more elegant and preferred way is to implement the operators as methods, as it is done everywhere in the standard library.

### Implement the operators as methods

Methods can be differentiated according to their receiver type. Therefore, instead of having to use different function names, we can simply define an  `Add()`  method for each type:

```go
func (a *sparseMatrix) Add(b Matrix) Matrix
func (a *denseMatrix) Add(b Matrix) Matrix
```

Each method returns a new object, which becomes the receiver of the next method call, so we can make chained

```go
expressions: m := m1.Mult(m2).Add(m3)
```

This is shorter and clearer. The correct implementation can again be selected at runtime based on a type-switch:

```go
func (a *sparseMatrix) Add(b Matrix) Matrix {
  switch b.(type) {
  case sparseMatrix:
    return addSparseToSparse(a.(sparseMatrix), b.(sparseMatrix))
  case denseMatrix:
    return addSparseToDense(a.(sparseMatrix), b.(denseMatrix))
  ...
  default:
    // unsupported arguments
  ...
  }
}
```

Again, this is easier than the nested type switch.

### Using an interface

When operating with the same methods on different types, the concept of creating a generalizing interface to implement this polymorphism should come to mind. We could, for example, define the interface  `Algebraic`:

```go
type Algebraic interface {
  Add(b Algebraic) Algebraic
  Min(b Algebraic) Algebraic
  Mult(b Algebraic) Algebraic
  ...
  Elements()
}
```

and define the methods  `Add()`,  `Min()`,  `Mult()`, … for our matrix types.

Each type that implements the  `Algebraic`  interface above will allow for method chaining. Each method implementation should use a type-switch to provide optimized implementations based on the parameter type. Additionally, a default case should be specified, which relies only on the methods in the interface:

```go
func (a *denseMatrix) Add(b Algebraic) Algebraic {
  switch b.(type) {
    case sparseMatrix:
      return addDenseToSparse(a, b.(sparseMatrix))
    default:
      for x in range b.Elements() ...
  ...
}
```

If a generic implementation cannot be implemented using only the methods in the interface, you are probably dealing with classes that are not similar enough, and this operator pattern should be abandoned. For example, it does not make sense to write  `a.Add(b)`  if  `a`  is a set, and  `b`  is a matrix; therefore, it will be difficult to implement a generic  `a.Add(b)`  in terms of set and matrix operators. In this case, split your package in two and provide separate  `AlgebraicSet`  and  `AlgebraicMatrix`  interfaces.



# Strings, Arrays and Slices

## 📝 Useful code snippets for strings

### Changing a character in a string

Strings are immutable, so in fact a new string is created here.

```go
str := "hello"
c := []rune(s)
c[0] = 'c'
s2 := string(c) // s2 == "cello"
```

### Taking a part (substring) of a string

```go
substr := str[n:m]
```

### Looping over a string with  `for`  or  `for-range`

```go
// gives only the bytes:
for i:=0; i < len(str); i++ {
  ... = str[i]
}
// gives the Unicode characters:
for ix, ch := range str {
  ...
}
```

### Finding number of bytes and characters in string

Number of bytes in a string  `str`:

```go
len(str)
```

Number of characters in a string  `str`:

The fastest way is:

```go
utf8.RuneCountInString(str)
```

An equivalent way is:

```go
len([]int(str))
```

### Concatenating strings

The fastest way is:

```go
// with a bytes.Buffer 
var buffer bytes.Buffer
var s string
buffer.WriteString(s)
fmt.Print(buffer.String(), "\n")
```

Other ways are:

```go
Strings.Join() // using Join function
str1 += str2 // using += operator
```

## 📝 Useful code snippets for arrays and slices

### Creation
To create an array:

```go
arr1 := new([len]type)
```

To create a slice:

```go
slice1 := make([]type, len)
```

### Initialization

To initialize an array:

```go
arr1 := [...]type{i1, i2, i3, i4, i5}
arrKeyValue := [len]type{i1: val1, i2: val2}
```

To initialize a slice:

```go
var slice1 []type = arr1[start:end]
```

### Cutting the last element of an array or slice line

```go
line = line[:len(line)-1]
```

### Looping over an array (or slice) arr with  `for`  or  `for-range`

```go
for i:=0; i < len(arr); i++ {
  ... = arr[i]
}

for ix, value := range arr {
  ...
}
```

# Structs, Interfaces and Maps

## 📝 Useful code snippets for structs

### Creation

```go
type struct1 struct {
  field1 type1
  field2 type2
  ...
}

ms := new(struct1)
```

### Initialization

```go
ms := &struct1{10, 15.5, "Chris"}
```

Capitalize the first letter of the struct name to make it visible outside its package. Often, it is better to define a factory function for the struct and force using that.

```go
ms := Newstruct1{10, 15.5, "Chris"}

func Newstruct1(n int, f float32, name string) *struct1 {
  return &struct1{n, f, name}
}
```

## 📝 Useful code snippets for interfaces

### Testing if a value implements an interface

```go
if v, ok := v.(Stringer); ok { // test if v implements Stringer
  fmt.Printf("implements String(): %s\n", v.String());
}
```

### A type classifier

```go
func classifier(items ...interface{}) {
  for i, x := range items {
    switch x.(type) {
    case bool: fmt.Printf("param #%d is a bool\n", i)
    case float64: fmt.Printf("param #%d is a float64\n", i)
    case int, int64: fmt.Printf("param #%d is an int\n", i)
    case nil: fmt.Printf("param #%d is nil\n", i)
    case string: fmt.Printf("param #%d is a string\n", i)

    default: fmt.Printf("param #%d's type is unknown\n", i)
    }
  }
}
```

## 📝 Useful code snippets for maps

### Creation

```go
map1 := make(map[keytype]valuetype)
```

### Initialization

```go
map1 := map[string]int{"one": 1, "two": 2}
```

### Looping over a map with  `for`or  `for-range`

```go
for key, value := range map1 {
  ...
}
```

### Testing if a key value exists in a map

```go
val1, isPresent = map1[key1]
```

This gives a value or zero-value for  `val1`,  _true_  or  _false_  for  `isPresent`.

### Deleting a key in a map

```go
delete(map1, key1)
```

# Functions and Files

## 📝 Useful code snippets for functions

### Recovering to stop a panic terminating sequence:

```go
func protect(g func()) {
  defer func() {
    log.Println("done") // Println executes normally even if there is a panic
    if x := recover(); x != nil {
      log.Printf("run time panic: %v", x)
    }
  }()
  log.Println("start")
  g()
}
```

## 📝 Useful code snippets for file

### Opening and reading a File

```go
file, err := os.Open("input.dat")
if err!= nil {
  fmt.Printf("An error occurred on opening the inputfile\n" +
  "Does the file exist?\n" +
  "Have you got acces to it?\n")
  return
}
defer file.Close()
iReader := bufio.NewReader(file)
for {
  str, err := iReader.ReadString('\n')
  if err!= nil {
    return // error or EOF
  }
  fmt.Printf("The input was: %s", str)
}
```

### Copying a file with a buffer

```go
func cat(f *file.File) {
  const NBUF = 512
  var buf [NBUF]byte
  for {
    switch nr, er := f.Read(buf[:]); true {
    case nr < 0:
      fmt.Fprintf(os.Stderr, "cat: error reading from %s: %s\n", f.String(),
      er.String())
      os.Exit(1)
    case nr == 0: // EOF
      return
    case nr > 0:
      if nw, ew := file.Stdout.Write(buf[0:nr]); nw != nr {
        fmt.Fprintf(os.Stderr, "cat: error writing from %s: %s\n", f.String(),
        ew.String())
      }
    }
  }
}
```

# Goroutines and Networking

A rule of thumb if you use  _parallelism_  to gain efficiency over serial computation:

_**"The amount of work done inside goroutines has to be much larger than the costs associated with creating goroutines and sending data back and forth between them."**_

-   **Using buffered channels for performance**: A buffered channel can easily double its throughput depending on the context, and the performance gain can be 10x or more. You can further try to optimize by adjusting the capacity of the channel.
-   **Limiting the number of items in a channel and packing them in arrays**: Channels become a bottleneck if you pass a lot of individual items through them. You can work around this by packing chunks of data into arrays and then unpacking on the other end. This can give a speed gain of a factor 10x.

## 📝 Useful code snippets for channels

### Creation

```go
ch := make(chan type, buf)
```

### Looping over a channel with a  `for–range`

```go
for v := range ch {
  // do something with v
}
```

### Testing if a channel is closed

```go
//read channel until it closes or error-condition
for {
  if input, open := <-ch; !open {
    break
  }
  fmt.Printf("%s ", input)
}
```

Or, use the  _looping over a channel method_  where the detection is automatic.

### Using a channel to let the main program wait

Using the semaphore pattern, you can let the main program wait until the goroutine completes:

```go
ch := make(chan int) // Allocate a channel

// Start something in a goroutine; when it completes, signal on the channel
go func() {
  // doSomething
  ch <- 1 // Send a signal; value does not matter
}()
doSomethingElseForAWhile()
<-ch // Wait for goroutine to finish; discard sent value
```

### Channel factory pattern

The function is a channel factory, and it starts a lambda function as goroutine, populating the channel:

```go
func pump() chan int {
  ch := make(chan int)
  go func() {
    for i := 0; ; i++ {
      ch <- i
    }
  }()
  return ch
}
```

### Stopping a goroutine

```go
runtime.Goexit()
```

### Simple timeout pattern

```go
timeout := make(chan bool, 1)
go func() {
  time.Sleep(1e9) // one second
  timeout <- true
}()
select {
  case <-ch:
    // a read from ch has occurred
  case <-timeout:
   // the read from ch has timed out
}
```

### Using an in- and out-channel instead of locking

```go
func Worker(in, out chan *Task) {
  for {
    t := <-in
    process(t)
    out <- t
  }
}
```

## 📝 Useful code snippets for networking applications

### Templating

-   Make, parse and validate a template:
    
    ```go
    var strTempl = template.Must(template.New("TName").Parse(strTemplateHTML))
    ```
    
-   Use the html filter to escape HTML special characters, when used in a web context:
    
    ```go
    {{html .}} or with a field FieldName {{ .FieldName |html }}
    ```
   
-   Use template-caching.

# Suggestions and Considerations

## General

### Stopping a program in case of error

```go
if err != nil {
  fmt.Printf("Program stopping with error %v", err)
  os.Exit(1)
}
```

or:

```go
if err != nil {
  panic("ERROR occurred: " + err.Error())
}
```

## Performance best practices and advice

-   Use  `[]rune`, if possible, instead of strings.
-   Use slices instead of arrays.
-   Use arrays or slices instead of a map where possible.
-   Use a  `for range`  over a slice if you only need the value and not the index; this is slightly faster than having to do a slice lookup for every element.
-   When the array is sparse (containing many 0 or nil-values), using a map can result in lower memory consumption.
-   Specify an initial capacity for maps.
-   When defining methods, use a pointer to a type (struct) as a receiver.
-   Use constants or flags to extract constant values from the code.
-   Use caching whenever possible when large amounts of memory are being allocated.
-   Use template caching.

## Memory considerations

For long-lived processes, deploy your Go applications on a  **64-bit**  OS. Due to the current nature of the GC, big chunks of memory are pre-allocated, which could bring a long-lived Go process on a 32-bit machine to crash due to memory depletion. Other causes of memory leaks could be:

-   You keep creating goroutines that don’t ever exit.
-   You keep appending to the same persistent slice (or writing to the same  `bytes.Buffer`) and never resetting it.
-   You keep adding items to a persistent map without ever removing them.

Here is a code snippet to monitor memory usage:

```go
memstats = new(runtime.MemStats)
runtime.ReadMemStats(memstats)
log.Printf("memstats before GC: bytes = %d footprint = %d",
memstats.HeapAlloc, memstats.Sys)
```

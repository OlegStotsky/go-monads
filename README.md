# go-monads
go-monads is a library that implements basic Haskell monads, based on Go 2 Generics Draft.

## Table of Contents

* [Maybe(T)](#maybet)
* [IO(T)](#iot)
* [Either(L, R)](#eitherl-r)
* [State(S, A)](#states-a)


## Maybe(T)
Maybe represents type that can either have value or be empty
### Create Maybe
```go
x := maybe.Return(int)(5)

or

val := new(int)
x := maybe.OfNullable(int)(val)

or

x := maybe.Empty(struct{})()
```

### Transform Maybe into new value
```go
x := maybe.Return(int)(5)
fmt.Println(Map(int, string)(x, fmt.Sprint))

or

func divideBy(x int, y int) Maybe(int) {
  if y == 0 { 
    return Empty(int)() 
  }
  return Return(int)(x/y)
}
four := maybe.Map(int, int)(divideBy(6, 3), func (x int) int { return x * 2 }) //four equals Just{obj: 4}
nothing := maybe.Map(int, int)(divideBy(6, 0), func (x int) int { return x * 2 }) //nothing equals Nothing{}

or

maybe.FlatMap(int, int)(divideBy(6, x), func (x int) { return divideBy(x, y) }) //x, y are some unknown integers, might be zeros
```

## IO(T)
IO encodes computation that potentially contains side effects as pure value
### Create IO
```go
//countNumberOfBytesInFile is an effectfull computation that returns number of bytes in file
func countNumberOfBytesInFile(f os.File) int {
  fi, _ := f.Stat()
  return fi.Size()
}

x := io.Return(int)(countNumberOfBytesIntFile(someFile))
```

### Transform IO into new value
```go
func printVal(type T)(val T) IO(T) {
  return Return(int)(func () { 
   fmt.Println(val)
   return val
 }) 
}

func main() {
  io.FlatMap(int, int)(countNumberOfBytesInFile(), printVal).UnsafePerformIO() 
}
```

## Either(L, R)
Either represents value that can be of one of two types

### Create Either
```go
x := either.AsRight(int)(5)
y := either.AsLeft(string)("xyz")
```

### Example of usage
For example, we want to check whether file contains substring 'go'

Some helper functions
```go
func toString(b []byte) string {
    return string(b)
}

func contains(text string) bool {
    return strings.Contains(text, "go")
}
```

Standard way
```go
func ContainsGo(reader io.Reader) (bool, error) {
    bytes, err:= ioutil.ReadAll(reader)
    if err != nil {
    return false, err
    }
    text := toString(bytes)
    return contains(text), nil
}
```
ReadAllE - function that call ReaderAll and wrap result to Either
```go
func ReadAllE(reader io.Reader) Either(error, []byte) 
```
Generic way
```go
func ContainsGoE(reader io.Reader) Either(error, bool) {
    return either.Map(error, string, bool)(either.Map(error, []byte, string)(ReadAllE(reader), toString), contains)	
}
```

## State(S, A)
Represent function func (S) (A, S), where S is state, A is result

### Create State
```go
x := state.Return(int, string)(5)
y := state.Put(int)(5)
```

### Transform State into new value
```go
func CountZeroes(numbs []int) int {
    counter := state.Put(int)(0)
    for _, x := range numbs {
        state := state.bindI(int, Unit, int)(counter, state.Get(int)())
        counter = state.FlatMap(int, int, int)(state, func(c int) { return state.Put(int)(c + 1) })
    }
    return counter.RunState(0)
}
```

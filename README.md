# go-monads
go-monads is a library that implements basic Haskell monads, based on Go 2 Generics Draft.

## Table of Contents

* [Maybe(T)](#maybet)
* [IO(T)](#iot)


## Maybe(T)
Maybe represents type that can either have value or be empty
### Create Maybe
```go
x := Return(int)(5)

or

val := new(int)
x := OfNullable(int)(val)

or

x := Empty(struct{})()
```

### Transform Maybe into new value
```go
x := Return(int)(5)
fmt.Println(Map(int, string)(x, fmt.Sprint))

or

func divideBy(x int, y int) Maybe(int) {
  if y == 0 { 
    return Empty(int)() 
  }
  return Return(int)(x/y)
}
four := Map(int, int)(divideBy(6, 3), func (x int) int { return x * 2 }) //four equals Just{obj: 4}
nothing := Map(int, int)(divideBy(6, 0), func (x int) int { return x * 2 }) //nothing equals Nothing{}

or

FlatMap(divideBy(6, x), func (x int) { return divideBy(x, y) }) //x, y are some unknown integers, might be zeros
```

## IO(T)
IO encodes computation that potentially contains side effects as pure value
### Create IO
```go
//countNumberOfBytesInFile is an effectfull computation that returns number of bytes in file
func countNumberOfBytesInFile() int {
  fi, _ := f.Stat()
  return fi.Size()
}

x := Return(int)(countNumberOfBytesIntFile)
```

### Transform IO into new value
```go
//
func printVal(type T)(val T) IO(T) {
  return Return(int)(func () { 
   fmt.Println(val)
   return val
 }) 
}

func main() {
  FlatMap(countNumberOfBytesInFile(), printVal).UnsafePerformIO() 
}
```
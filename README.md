# go-monads
go-monads is a library that implements basic Haskell monads, based on Go 2 Generics Draft.

## Table of Contents

* [Maybe[T]](#maybet)
* [IO[T]](#iot)
* [Either[L, R]](#eitherl-r)
* [State[S, A]](#states-a)


## Maybe[T]
Maybe represents type that can either have value or be empty
### Create Maybe
```go
m1 := maybe.Return(5)

or

val := new(int)
m2 := maybe.OfNullable[int](val)

or

m3 := maybe.Empty[any]()
```

### Transform Maybe into new value
```go
x := maybe.Return[int](5)
fmt.Println(maybe.Map[int, string](x, strconv.Itoa).Get())

or


func divideBy(x int, y int) maybe.Maybe[int] {
    if y == 0 {
        return maybe.Empty[int]()
    }
    return maybe.Return(x / y)
}

four := maybe.Map(divideBy(6, 3), func(x int) int { return x * 2 })    //four equals Just{obj: 4}
nothing := maybe.Map(divideBy(6, 0), func(x int) int { return x * 2 }) //nothing equals Nothing{}

or

x := 0
y := 7
m5 := maybe.FlatMap[int, int](divideBy(6, x), func(x int) maybe.Maybe[int] { return divideBy(x, y) }) //x, y are some unknown integers, might be zeros
```

## IO[T]
IO encodes computation that potentially contains side effects as pure value
### Create IO
```go
//countNumberOfBytesInFile is an effectfull computation that returns number of bytes in file
func countNumberOfBytesInFile(f os.File) int {
  fi, _ := f.Stat()
  return fi.Size()
}

f, err := os.CreateTemp("./", "ab*")
if err != nil {
    log.Fatalln(err.Error())
}
defer os.Remove(f.Name())
f.WriteString("hello world")

x := io.Return[int64](func() int64 { return countNumberOfBytesInFile(f) })
fmt.Println(x.UnsafePerformIO())
```

### Transform IO into new value
```go
func printVal[T any](val T) io.IO[util.Unit] {
    return io.Return[util.Unit](func() util.Unit {
        fmt.Println(val)
        return util.UnitVar
    })
}

func main() {
    x2 := io.FlatMap(x, printVal[int64])
    fmt.Println(x2.UnsafePerformIO())
}
```

## Either(L, R)
Either represents value that can be of one of two types

### Create Either
```go
x := either.AsRight(5)
y := either.AsLeft("xyz")
fmt.Println(either.ToMaybe(x).Get()) // prints 5, nil
fmt.Println(either.ToMaybe(y).Get()) // prints <nil> Trying to get from Nothing
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
func ReadAllE(reader io.Reader) either.Either[error, []byte] {
	x := either.FromErrorable[[]byte](ioutil.ReadAll(reader))
	return x
}
```
Generic way
```go
func ContainsGoE(reader io.Reader) either.Either[error, bool] {
    return either.Map(either.Map(ReadAllE(reader), toString), contains)
}
```

## State(S, A)
Represent function func (S) (A, S), where S is state, A is result

### Create State
```go
x := state.Return[any, int](5)
y := state.Put(100)
fmt.Println(x.RunState(nil)) // prints 5
fmt.Println(y.RunState(6))   // prints 100
```

### Transform State into new value
```go
func CountZeroes(ints []int) int {
    counter := state.Put[int](0)
    for _, x := range ints {
        if x == 0 {
            s := state.BindI[int, util.Unit, int](counter, state.Get[int]())
            counter = state.FlatMap[int, int, util.Unit](s, func(c int) state.State[int, util.Unit] { return state.Put[int](c + 1) })
        }
    }
    _, s := counter.RunState(0)
    return s
}
```

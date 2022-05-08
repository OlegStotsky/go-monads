package main

import (
	"fmt"
	"github.com/OlegStotsky/go-monads/maybe"
	"strconv"
)

func main() {
	m1 := maybe.Return(5)
	fmt.Println(m1.Get())

	val := new(int)
	m2 := maybe.OfNullable[int](val)
	fmt.Println(m2.Get())

	m3 := maybe.Empty[any]()
	fmt.Println(m3.Get())

	m4 := maybe.Return[int](5)
	fmt.Println(maybe.Map[int, string](m4, strconv.Itoa).Get())

	four := maybe.Map(divideBy(6, 3), func(x int) int { return x * 2 })    //four equals Just{obj: 4}
	nothing := maybe.Map(divideBy(6, 0), func(x int) int { return x * 2 }) //nothing equals Nothing{}
	fmt.Println(four.Get())
	fmt.Println(nothing.Get())

	x := 0
	y := 7
	m5 := maybe.FlatMap[int, int](divideBy(6, x), func(x int) maybe.Maybe[int] { return divideBy(x, y) }) //x, y are some unknown integers, might be zeros
	fmt.Println(m5.Get())
}

func divideBy(x int, y int) maybe.Maybe[int] {
	if y == 0 {
		return maybe.Empty[int]()
	}
	return maybe.Return(x / y)
}

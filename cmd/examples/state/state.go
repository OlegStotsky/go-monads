package main

import (
	"fmt"
	"github.com/OlegStotsky/go-monads/state"
	"github.com/OlegStotsky/go-monads/util"
)

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

func main() {
	x := state.Return[any, int](5)
	y := state.Put(100)
	fmt.Println(x.RunState(nil)) // prints 5
	fmt.Println(y.RunState(6))   // prints 100

	fmt.Println(CountZeroes([]int{5, 4, 0, 0, 1, 0})) // prints 3
}

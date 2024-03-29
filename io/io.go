package io

import (
	"github.com/OlegStotsky/go-monads/util"
)

//IO encodes computation that potentially contains side effects as pure value
//UnsafePerformIO runs computation. Should only be called in main function
type IO[U any] interface {
	UnsafePerformIO() U
}

type ioImpl[U any] struct {
	unsafeRun util.Func0[U]
}

func (i ioImpl[U]) UnsafePerformIO() U {
	return i.unsafeRun()
}

//Constructs IO from function
func Return[U any](f util.Func0[U]) IO[U] {
	return ioImpl[U]{unsafeRun: f}
}

//Map takes IO and function f and returns new IO that when run will perform computations sequentially
func Map[U, V any](m IO[U], f util.Func[U, V]) IO[V] {
	return ioImpl[V]{unsafeRun: func() V { return f(m.UnsafePerformIO()) }}
}

//FlatMap creates composite action that represents action m followed by action f
func FlatMap[U, V any](m IO[U], f util.Func[U, IO[V]]) IO[V] {
	return f(m.UnsafePerformIO())
}

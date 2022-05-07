package either

import (
	"github.com/OlegStotsky/go-monads/maybe"
	"github.com/OlegStotsky/go-monads/util"
)

//Either represents value that can be of one of two types
type Either[L, R any] interface{}

//Right represents right value of Either
type Right[R any] struct {
	obj R
}

//Left represents left value of Either
type Left[L any] struct {
	obj L
}

//AsRight creates Right with given value
func AsRight[R any](x R) Either[any, R] {
	return Right[R]{obj: x}
}

//AsLeft create Left with given value
func AsLeft[L any](x L) Either[L, any] {
	return Left[L]{obj: x}
}

//Return Right containing the result of applying f to m if it's not Left. Otherwise returns unchanged Left
func Map[L, R, V any](e Either[L, R], f util.Func[R, V]) Either[L, V] {
	switch e.(type) {
	case Right[R]:
		return Right[V]{obj: f(e.(Right[R]).obj)}
	}
	return e
}

//If this is Right returns result of applying f. Otherwise returns unchanged Left
func FlatMap[L, R, V any](e Either[L, R], f util.Func[R, Either[L, V]]) Either[L, V] {
	switch e.(type) {
	case Right[R]:
		return f(e.(Right[R]).obj)
	}
	return e
}

//Returns contained value if this is Right. Otherwise returns given value
func OrElse[L, R any](e Either[L, R], other R) R {
	switch e.(type) {
	case Right[R]:
		return e.(Right[R]).obj
	}
	return other
}

//If this is Right returns Left with same value or vice versa
func Swap[L, R any](e Either[L, R]) Either[R, L] {
	switch e.(type) {
	case Right[R]:
		return AsLeft(e.(Right[R]).obj)
	}
	return AsRight(e.(Left[L]).obj)
}

//Converts Right to Just, Left to Nothing
func ToMaybe[L, R any](e Either[L, R]) maybe.Maybe[R] {
	switch e.(type) {
	case Right[R]:
		return maybe.Return(e.(Right[R]).obj)
	}
	return maybe.Nothing[R]{}
}

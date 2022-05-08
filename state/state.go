package state

import "github.com/OlegStotsky/go-monads/util"

//Represent function func (S) (A, S), where S is state, A is result
//RunState run state function with given state
type State[S, A any] interface {
	RunState(S) (A, S)
}

type stateImpl[S, A any] struct {
	f func(S) (A, S)
}

func (s stateImpl[S, A]) RunState(state S) (A, S) {
	return s.f(state)
}

//Create State with given value
func Return[S, A any](x A) State[S, A] {
	return stateImpl[S, A]{f: func(s S) (A, S) { return x, s }}
}

//Apply given function to result of handling given state
func Map[S, A, B any](s State[S, A], mapF util.Func[A, B]) State[S, B] {
	return stateImpl[S, B]{f: func(state S) (B, S) {
		a, newState := s.RunState(state)
		return mapF(a), newState
	}}
}

//Apply given function to result of handling given state and produce new state
func FlatMap[S, A, B any](s State[S, A], mapF util.Func[A, State[S, B]]) State[S, B] {
	return stateImpl[S, B]{f: func(state S) (B, S) {
		a, newState := s.RunState(state)
		result := mapF(a)
		return result.RunState(newState)
	}}
}

//Give state as value of State
func Get[S any]() State[S, S] {
	return stateImpl[S, S]{f: func(state S) (S, S) {
		return state, state
	}}
}

//Put new state to State
func Put[S any](newState S) State[S, util.Unit] {
	return stateImpl[S, util.Unit]{f: func(S) (util.Unit, S) { return struct{}{}, newState }}
}

//Analogue of flatMap with ignoring value of base State
func BindI[S, A, B any](s State[S, A], s2 State[S, B]) State[S, B] {
	return stateImpl[S, B]{f: func(state S) (B, S) {
		_, newState := s.RunState(state)
		return s2.RunState(newState)
	}}
}

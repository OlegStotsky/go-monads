package maybe

import (
	"errors"
	"github.com/OlegStotsky/go-monads/util"
)

//Maybe represents type that can either have value or be empty
//IsPresent tells whether instance holds value, in which case it returns true.
//Filter returns this Maybe if it's not empty and satisfies predicate. Otherwise it returns Nothing
//OrElse returns this Maybe if it's not empty, otherwise it returns other
//Get returns value of this Maybe, or error
type Maybe[U any] interface {
	IsPresent() bool
	Filter(p util.Predicate[U]) Maybe[U]
	OrElse(other U) U
	Get() (U, error)
}

//Just represents non empty case of Maybe
type Just[U any] struct {
	obj U
}

//Nothing represents empty case of Maybe
type Nothing[U any] struct {
}

//Return creates Maybe from value
func Return[U any](x U) Maybe[U] {
	return Just[U]{obj: x}
}

//OfNullable creates Maybe from value that can be nil
//In case of nil it creates Nothing
func OfNullable[U any](x *U) Maybe[U] {
	if x == nil {
		return Nothing[U]{}
	}

	return Just[U]{obj: *x}
}

//Empty creates Nothing
func Empty[U any]() Maybe[U] {
	return Nothing[U]{}
}

func (j Just[U]) IsPresent() bool {
	return true
}

func (n Nothing[U]) IsPresent() bool {
	return false
}

func (j Just[U]) Filter(p util.Predicate[U]) Maybe[U] {
	if p(j.obj) {
		return j
	}
	return Nothing[U]{}
}

//Filter on Nothing returns Nothing
func (n Nothing[U]) Filter(p util.Predicate[U]) Maybe[U] {
	return n
}

func (j Just[U]) OrElse(other U) U {
	return j.obj
}

func (n Nothing[U]) OrElse(other U) U {
	return other
}

func (j Just[U]) Get() (U, error) {
	return j.obj, nil
}

var NoElementError = errors.New("Trying to get from Nothing")

func (n Nothing[U]) Get() (U, error) {
	return *new(U), NoElementError
}

//Returns Just containing the result of applying f to m if it's non empty. Otherwise returns Nothing
func Map[U, V any](m Maybe[U], f util.Func[U, V]) Maybe[V] {
	switch m.(type) {
	case Just[U]:
		return Just[V]{obj: f(m.(Just[U]).obj)}
	}
	return Nothing[V]{}
}

//Same as Map but function must return Maybe
func FlatMap[U, V any](m Maybe[U], f util.Func[U, Maybe[V]]) Maybe[V] {
	switch m.(type) {
	case Just[U]:
		return f(m.(Just[U]).obj)
	}
	return Nothing[V]{}
}

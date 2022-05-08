package main

import (
	"fmt"
	"github.com/OlegStotsky/go-monads/either"
	"io"
	"io/ioutil"
	"strings"
)

func toString(b []byte) string {
	return string(b)
}

func contains(text string) bool {
	return strings.Contains(text, "go")
}

func containsGo(reader io.Reader) (bool, error) {
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return false, err
	}
	text := toString(bytes)
	return contains(text), nil
}

func ReadAllE(reader io.Reader) either.Either[error, []byte] {
	x := either.FromErrorable[[]byte](ioutil.ReadAll(reader))
	return x
}

func ContainsGoE(reader io.Reader) either.Either[error, bool] {
	return either.Map(either.Map(ReadAllE(reader), toString), contains)
}

func main() {
	fmt.Println("first example")
	x := either.AsRight(5)
	y := either.AsLeft("xyz")
	fmt.Println(either.ToMaybe(x).Get()) // prints 5, nil
	fmt.Println(either.ToMaybe(y).Get()) // prints <nil> Trying to get from Nothing

	fmt.Println("second example")
	r := strings.NewReader("hello world go")
	fmt.Println(containsGo(r))
	r.Reset("hello world go")
	fmt.Println(either.ToMaybe(ContainsGoE(r)).Get())
}

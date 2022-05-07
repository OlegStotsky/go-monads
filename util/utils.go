package util

type Func[U, V any] func(U) V
type Func0[V any] func() V
type Predicate[U any] func(U) bool
type Unit struct{}

var UnitVar = struct{}{}

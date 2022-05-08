package main

import (
	"fmt"
	"github.com/OlegStotsky/go-monads/io"
	"github.com/OlegStotsky/go-monads/util"
	"log"
	"os"
)

//countNumberOfBytesInFile is an effectfull computation that returns number of bytes in file
func countNumberOfBytesInFile(f *os.File) int64 {
	fi, _ := f.Stat()
	return fi.Size()
}

func printVal[T any](val T) io.IO[util.Unit] {
	return io.Return[util.Unit](func() util.Unit {
		fmt.Println(val)
		return util.UnitVar
	})
}

func main() {
	f, err := os.CreateTemp("./", "ab*")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer os.Remove(f.Name())
	f.WriteString("hello world")

	x := io.Return[int64](func() int64 { return countNumberOfBytesInFile(f) })
	fmt.Println(x.UnsafePerformIO())

	x2 := io.FlatMap(x, printVal[int64])
	fmt.Println(x2.UnsafePerformIO())
}

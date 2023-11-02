package main

import (
	_ "embed"
	"fmt"
)

//go:embed foo.txt
var Foo string

func main() {

	fmt.Println(Foo)
}

func GetFoo() string {
	return Foo
}

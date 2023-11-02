package foo

import (
	_ "embed"
)

//go:embed foo.txt
var Foo string

func GetFoo() string {
	return Foo
}

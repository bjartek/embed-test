package main

import (
	_ "embed"
	"fmt"

	"github.com/bjartek/embed-test/foo"
)

func main() {

	fmt.Println(foo.GetFoo())
}

package main

import (
	_ "embed"
	"fmt"
)

//go:embed foo.txt
var foo string

func main() {

	fmt.Println(foo)
}
